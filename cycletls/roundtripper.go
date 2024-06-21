package cycletls

import (
	"context"
	"errors"
	"fmt"
	"net"

	fhttp "github.com/aurorax-neo/tls_client_httpi/cycletls_fhttp"
	"github.com/aurorax-neo/tls_client_httpi/cycletls_fhttp/http2"
	utls "github.com/refraction-networking/utls"
	"golang.org/x/net/proxy"
	"strings"
	"sync"
)

var errProtocolNegotiated = errors.New("protocol negotiated")

type roundTripper struct {
	sync.Mutex
	// fix typing
	JA3       string
	UserAgent string

	InsecureSkipVerify bool
	Cookies            []Cookie
	cachedConnections  map[string]net.Conn
	cachedTransports   map[string]fhttp.RoundTripper

	dialer     proxy.ContextDialer
	forceHTTP1 bool
}

func (rt *roundTripper) RoundTrip(req *fhttp.Request) (*fhttp.Response, error) {
	// Fix this later for proper cookie parsing
	for _, properties := range rt.Cookies {
		req.AddCookie(&fhttp.Cookie{
			Name:       properties.Name,
			Value:      properties.Value,
			Path:       properties.Path,
			Domain:     properties.Domain,
			Expires:    properties.JSONExpires.Time, //TODO: scuffed af
			RawExpires: properties.RawExpires,
			MaxAge:     properties.MaxAge,
			HttpOnly:   properties.HTTPOnly,
			Secure:     properties.Secure,
			Raw:        properties.Raw,
			Unparsed:   properties.Unparsed,
		})
	}
	req.Header.Set("User-Agent", rt.UserAgent)
	addr := rt.getDialTLSAddr(req)
	if _, ok := rt.cachedTransports[addr]; !ok {
		if err := rt.getTransport(req, addr); err != nil {
			return nil, err
		}
	}
	return rt.cachedTransports[addr].RoundTrip(req)
}

func (rt *roundTripper) getTransport(req *fhttp.Request, addr string) error {
	switch strings.ToLower(req.URL.Scheme) {
	case "fhttp":
		rt.cachedTransports[addr] = &fhttp.Transport{DialContext: rt.dialer.DialContext, DisableKeepAlives: true}
		return nil
	case "https":
	default:
		return fmt.Errorf("invalid URL scheme: [%v]", req.URL.Scheme)
	}

	_, err := rt.dialTLS(req.Context(), "tcp", addr)
	switch {
	case errors.Is(err, errProtocolNegotiated):
	case err == nil:
		// Should never happen.
		panic("dialTLS returned no error when determining cachedTransports")
	default:
		return err
	}

	return nil
}

func (rt *roundTripper) dialTLS(ctx context.Context, network, addr string) (net.Conn, error) {
	rt.Lock()
	defer rt.Unlock()

	// If we have the connection from when we determined the HTTPS
	// cachedTransports to use, return that.
	if conn := rt.cachedConnections[addr]; conn != nil {
		return conn, nil
	}
	rawConn, err := rt.dialer.DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}

	var host string
	if host, _, err = net.SplitHostPort(addr); err != nil {
		host = addr
	}
	//////////////////

	spec, err := StringToSpec(rt.JA3, rt.UserAgent, rt.forceHTTP1)
	if err != nil {
		return nil, err
	}

	conn := utls.UClient(rawConn, &utls.Config{ServerName: host, OmitEmptyPsk: true, InsecureSkipVerify: rt.InsecureSkipVerify}, // MinVersion:         tls.VersionTLS10,
		// MaxVersion:         tls.VersionTLS13,

		utls.HelloCustom)

	if err := conn.ApplyPreset(spec); err != nil {
		return nil, err
	}

	if err = conn.Handshake(); err != nil {
		_ = conn.Close()

		if err.Error() == "tls: CurvePreferences includes unsupported curve" {
			//fix this
			return nil, fmt.Errorf("conn.Handshake() error for tls 1.3 (please retry request): %+v", err)
		}
		return nil, fmt.Errorf("uTlsConn.Handshake() error: %+v", err)
	}

	if rt.cachedTransports[addr] != nil {
		return conn, nil
	}

	// No fhttp.Transport constructed yet, create one based on the results
	// of ALPN.
	switch conn.ConnectionState().NegotiatedProtocol {
	case http2.NextProtoTLS:
		parsedUserAgent := parseUserAgent(rt.UserAgent)

		t2 := http2.Transport{
			DialTLS:     rt.dialTLSHTTP2,
			PushHandler: &http2.DefaultPushHandler{},
			Navigator:   parsedUserAgent.UserAgent,
		}
		rt.cachedTransports[addr] = &t2
	default:
		// Assume the remote peer is speaking HTTP 1.x + TLS.
		rt.cachedTransports[addr] = &fhttp.Transport{DialTLSContext: rt.dialTLS, DisableKeepAlives: true}

	}

	// Stash the connection just established for use servicing the
	// actual request (should be near-immediate).
	rt.cachedConnections[addr] = conn

	return nil, errProtocolNegotiated
}

func (rt *roundTripper) dialTLSHTTP2(network, addr string, _ *utls.Config) (net.Conn, error) {
	return rt.dialTLS(context.Background(), network, addr)
}

func (rt *roundTripper) getDialTLSAddr(req *fhttp.Request) string {
	host, port, err := net.SplitHostPort(req.URL.Host)
	if err == nil {
		return net.JoinHostPort(host, port)
	}
	return net.JoinHostPort(req.URL.Host, "443") // we can assume port is 443 at this point
}

func (rt *roundTripper) CloseIdleConnections() {
	for addr, conn := range rt.cachedConnections {
		_ = conn.Close()
		delete(rt.cachedConnections, addr)
	}
}

func newRoundTripper(browser Browser, dialer ...proxy.ContextDialer) fhttp.RoundTripper {
	if len(dialer) > 0 {

		return &roundTripper{
			dialer:             dialer[0],
			JA3:                browser.JA3,
			UserAgent:          browser.UserAgent,
			Cookies:            browser.Cookies,
			cachedTransports:   make(map[string]fhttp.RoundTripper),
			cachedConnections:  make(map[string]net.Conn),
			InsecureSkipVerify: browser.InsecureSkipVerify,
			forceHTTP1:         browser.forceHTTP1,
		}
	}

	return &roundTripper{
		dialer:             proxy.Direct,
		JA3:                browser.JA3,
		UserAgent:          browser.UserAgent,
		Cookies:            browser.Cookies,
		cachedTransports:   make(map[string]fhttp.RoundTripper),
		cachedConnections:  make(map[string]net.Conn),
		InsecureSkipVerify: browser.InsecureSkipVerify,
		forceHTTP1:         browser.forceHTTP1,
	}
}
