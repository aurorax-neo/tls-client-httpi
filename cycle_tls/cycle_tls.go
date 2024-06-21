package cycle_tls

import (
	"github.com/aurorax-neo/tls_client_httpi"
	"github.com/aurorax-neo/tls_client_httpi/cycletls"
	fhttp "github.com/aurorax-neo/tls_client_httpi/cycletls_fhttp"
	"golang.org/x/net/context"
	"golang.org/x/net/proxy"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

type CycleTls struct {
	Client     *fhttp.Client
	ReqBefore  handler
	ja3        string
	ua         string
	timeOutSec int
}

type handler func(req *fhttp.Request) error

func NewClient(ja3 string, ua string, timeOutSec int) *CycleTls {
	client := &fhttp.Client{
		Transport: cycletls.NewTransport(ja3, ua),
		Timeout:   time.Duration(timeOutSec) * time.Second,
	}
	return &CycleTls{
		Client:     client,
		ja3:        ja3,
		ua:         ua,
		timeOutSec: timeOutSec,
	}
}

func DefaultClient() *CycleTls {
	ja3 := "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,17513-43-5-35-45-27-10-18-65281-0-23-16-51-11-13-21,29-23-24,0"
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"
	return NewClient(ja3, ua, 30)
}

func convertResponse(resp *fhttp.Response) *http.Response {
	response := &http.Response{
		Status:           resp.Status,
		StatusCode:       resp.StatusCode,
		Proto:            resp.Proto,
		ProtoMajor:       resp.ProtoMajor,
		ProtoMinor:       resp.ProtoMinor,
		Header:           http.Header(resp.Header),
		Body:             resp.Body,
		ContentLength:    resp.ContentLength,
		TransferEncoding: resp.TransferEncoding,
		Close:            resp.Close,
		Uncompressed:     resp.Uncompressed,
		Trailer:          http.Header(resp.Trailer),
	}
	return response
}

func (CT *CycleTls) handleHeaders(req *fhttp.Request, headers tls_client_httpi.Headers) {
	if headers == nil {
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func (CT *CycleTls) handleCookies(req *fhttp.Request, cookies tls_client_httpi.Cookies) {
	if cookies == nil {
		return
	}
	for _, c := range cookies {
		req.AddCookie(&fhttp.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			Expires:    c.Expires,
			RawExpires: c.RawExpires,
			MaxAge:     c.MaxAge,
			Secure:     c.Secure,
			HttpOnly:   c.HttpOnly,
			SameSite:   fhttp.SameSite(c.SameSite),
			Raw:        c.Raw,
			Unparsed:   c.Unparsed,
		})
	}
}

func (CT *CycleTls) Request(method tls_client_httpi.Method, rawURL string, headers tls_client_httpi.Headers, cookies tls_client_httpi.Cookies, body io.Reader) (*http.Response, error) {
	req, err := fhttp.NewRequest(string(method), rawURL, body)
	if err != nil {
		return nil, err
	}
	CT.handleHeaders(req, headers)
	CT.handleCookies(req, cookies)
	if CT.ReqBefore != nil {
		if err := CT.ReqBefore(req); err != nil {
			return nil, err
		}
	}
	do, err := CT.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return convertResponse(do), nil
}

func (CT *CycleTls) SetProxy(rawUrl string) error {
	// URL
	URL, err := url.Parse(rawUrl)
	if err != nil {
		return err
	}
	// 创建一个 Dialer
	dialer, err := proxy.FromURL(URL, proxy.Direct)
	if err != nil {
		return err
	}
	// 将 Dialer 转换为 ContextDialer
	contextDialer := &contextDialerAdapter{dialer}
	CT.Client = &fhttp.Client{
		Transport: cycletls.NewTransportWithProxy(CT.ja3, CT.ua, contextDialer),
		Timeout:   time.Duration(CT.timeOutSec) * time.Second,
	}
	return nil
}

func (CT *CycleTls) SetCookies(rawUrl string, cookies tls_client_httpi.Cookies) {
	if cookies == nil {
		return
	}
	u, err := url.Parse(rawUrl)
	if err != nil {
		return
	}
	var fCookies []*fhttp.Cookie
	for _, c := range cookies {
		fCookies = append(fCookies, &fhttp.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			Expires:    c.Expires,
			RawExpires: c.RawExpires,
			MaxAge:     c.MaxAge,
			Secure:     c.Secure,
			HttpOnly:   c.HttpOnly,
			SameSite:   fhttp.SameSite(c.SameSite),
			Raw:        c.Raw,
			Unparsed:   c.Unparsed,
		})
	}
	CT.Client.Jar.SetCookies(u, fCookies)
}

func (CT *CycleTls) GetCookies(rawUrl string) tls_client_httpi.Cookies {
	currUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil
	}

	var cookies tls_client_httpi.Cookies
	for _, c := range CT.Client.Jar.Cookies(currUrl) {
		cookies = append(cookies, &http.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			Expires:    c.Expires,
			RawExpires: c.RawExpires,
			MaxAge:     c.MaxAge,
			Secure:     c.Secure,
			HttpOnly:   c.HttpOnly,
			SameSite:   http.SameSite(c.SameSite),
		})
	}
	return cookies
}

// contextDialerAdapter 适配器将 proxy.Dialer 转换为 proxy.ContextDialer
type contextDialerAdapter struct {
	proxy.Dialer
}

func (d *contextDialerAdapter) DialContext(_ context.Context, network, addr string) (net.Conn, error) {
	return d.Dial(network, addr)
}
