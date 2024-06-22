package TCHI_test

import (
	"fmt"
	"github.com/aurorax-neo/tls_client_httpi"
	"github.com/aurorax-neo/tls_client_httpi/TCHUtil"
	"github.com/aurorax-neo/tls_client_httpi/tls_client"
	"github.com/bogdanfinn/fhttp/http2"
	"github.com/bogdanfinn/tls-client/profiles"
	tls "github.com/bogdanfinn/utls"
	"testing"
)

func TestGetReq(t *testing.T) {
	c := tls_client.NewClient(tls_client.NewClientOptions(30, profiles.Chrome_124))
	response, err := c.Request("GET", "https://tls.browserleaks.com/json", nil, nil, nil)
	if err != nil {
		return
	}
	TCHUtil.OutHttpResponse(response)
}

func TestGetProxy(t *testing.T) {
	c := tls_client.DefaultClient()
	c.SetProxy("http://127.0.0.1:7890")
	response, err := c.Request("GET", "https://www.ip.cn/api/index?ip&type=0", nil, nil, nil)
	if err != nil {
		return
	}

	fmt.Println("c")
	TCHUtil.OutHttpResponse(response)

}

func TestGetTls(t *testing.T) {
	// clientHelloId tls.ClientHelloID, settings map[http2.SettingID]uint32, settingsOrder []http2.SettingID, pseudoHeaderOrder []string, connectionFlow uint32, priorities []http2.Priority, headerPriority *http2.PriorityParam
	clientHelloId := tls.ClientHelloID{
		Client:               "Edge",
		RandomExtensionOrder: false,
		Version:              "117",
		Seed:                 nil,
		Weights:              nil,
		SpecFactory: func() (tls.ClientHelloSpec, error) {
			return tls.ClientHelloSpec{
				CipherSuites: []uint16{
					0xbaba,
					0x1301,
					0x1302,
					0x1303,
					0xc02b,
					0xc02f,
					0xc02c,
					0xc030,
					0xcca9,
					0xcca8,
					0xc013,
					0xc014,
					0x009c,
					0x009d,
					0x002f,
					0x0035,
				},
				CompressionMethods: []uint8{
					0,
				},
				Extensions: []tls.TLSExtension{
					&tls.UtlsGREASEExtension{},
					&tls.UtlsCompressCertExtension{Algorithms: []tls.CertCompressionAlgo{
						tls.CertCompressionBrotli,
					}},
					&tls.SCTExtension{},
					&tls.ExtendedMasterSecretExtension{},
					&tls.ApplicationSettingsExtension{SupportedProtocols: []string{"h2"}},
					&tls.ALPNExtension{AlpnProtocols: []string{"h2", "http/1.1"}},
					&tls.SupportedVersionsExtension{Versions: []uint16{
						tls.GREASE_PLACEHOLDER,
						tls.VersionTLS13,
						tls.VersionTLS12,
					}},
					&tls.SignatureAlgorithmsExtension{SupportedSignatureAlgorithms: []tls.SignatureScheme{
						tls.ECDSAWithP256AndSHA256,
						tls.PSSWithSHA256,
						tls.PKCS1WithSHA256,
						tls.ECDSAWithP384AndSHA384,
						tls.PSSWithSHA384,
						tls.PKCS1WithSHA384,
						tls.PSSWithSHA512,
						tls.PKCS1WithSHA512,
					}},
					&tls.SupportedPointsExtension{SupportedPoints: []byte{
						tls.PointFormatUncompressed,
					}},
					&tls.SNIExtension{},
					&tls.SessionTicketExtension{},
					&tls.SupportedCurvesExtension{Curves: []tls.CurveID{
						tls.GREASE_PLACEHOLDER,
						tls.X25519Kyber768Draft00,
						tls.X25519,
						tls.CurveP256,
						tls.CurveP384,
					}},
					tls.BoringGREASEECH(),
					&tls.StatusRequestExtension{},
					&tls.RenegotiationInfoExtension{Renegotiation: tls.RenegotiateOnceAsClient},
					&tls.PSKKeyExchangeModesExtension{Modes: []uint8{
						tls.PskModeDHE,
					}},
					&tls.KeyShareExtension{KeyShares: []tls.KeyShare{
						{Group: tls.CurveID(tls.GREASE_PLACEHOLDER), Data: []byte{0}},
						{Group: tls.X25519Kyber768Draft00},
						{Group: tls.X25519},
					}},
					&tls.UtlsGREASEExtension{},
				},
			}, nil
		},
	}
	settings := map[http2.SettingID]uint32{
		http2.SettingHeaderTableSize:   65536,
		http2.SettingEnablePush:        0,
		http2.SettingInitialWindowSize: 6291456,
		http2.SettingMaxHeaderListSize: 262144,
	}
	settingsOrder := []http2.SettingID{
		http2.SettingHeaderTableSize,
		http2.SettingEnablePush,
		http2.SettingInitialWindowSize,
		http2.SettingMaxHeaderListSize,
	}
	pseudoHeaderOrder := []string{
		":method",
		":authority",
		":scheme",
		":path",
	}
	connectionFlow := uint32(15663105)
	pro := profiles.NewClientProfile(clientHelloId, settings, settingsOrder, pseudoHeaderOrder, connectionFlow, nil, nil)
	cg := tls_client.NewClientOptions(30, pro)
	c := tls_client.NewClient(cg)
	response, err := c.Request("GET", "https://tls.browserleaks.com/json", nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	TCHUtil.OutHttpResponse(response)

	headers := tls_client_httpi.Headers{}
	headers.Set("accept", "*/*")
	headers.Set("accept-language", "zh-CN,zh;q=0.9,zh-Hans;q=0.8,en;q=0.7")
	headers.Set("oai-language", "en-US")
	//headers.Set("origin", common.GetOrigin(url))
	//headers.Set("referer", common.GetOrigin(url))
	headers.Set("sec-ch-ua", `"Microsoft Edge";v="123", "Not:A-Brand";v="8", "Chromium";v="123"`)
	headers.Set("sec-ch-ua-mobile", "?0")
	headers.Set("sec-ch-ua-platform", `"Windows"`)
	headers.Set("sec-fetch-dest", "empty")
	headers.Set("sec-fetch-mode", "cors")
	headers.Set("sec-fetch-site", "same-origin")
	headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36 Edg/117.0.2045.31")
	headers.Set("Connection", "close")
	c.SetProxy("http://127.0.0.1:7890")
	response, err = c.Request(tls_client_httpi.GET, "https://chatgpt.com", headers, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	TCHUtil.OutHttpResponse(response)

}
