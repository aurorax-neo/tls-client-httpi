package tls_client

import (
	"github.com/bogdanfinn/fhttp/http2"
	tlsclient "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	bfutls "github.com/bogdanfinn/utls"
)

func Edge117() profiles.ClientProfile {
	ja3String := "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,13-10-51-18-65281-27-0-23-35-43-16-11-5-45-17513-21,29-23-24,0"

	// find the following in tls.go
	_ = tlsclient.H2SettingsMap
	_ = bfutls.PKCS1WithSHA256

	supportedSignatureAlgorithms := []string{
		"ECDSAWithP256AndSHA256", // 0x0403
		"PSSWithSHA256",          // 0x0804
		"PKCS1WithSHA256",        // 0x0401
		"ECDSAWithP384AndSHA384", // 0x0503
		"PSSWithSHA384",          // 0x0805
		"PKCS1WithSHA384",        // 0x0501
		"PSSWithSHA512",          // 0x0806
		"PKCS1WithSHA512",        // 0x0601
	}

	supportedDelegatedCredentialsAlgorithms := []string{
		"ECDSAWithP256AndSHA256", // 0x0403
		"PSSWithSHA256",          // 0x0804
		"PKCS1WithSHA256",        // 0x0401
		"ECDSAWithP384AndSHA384", // 0x0503
		"PSSWithSHA384",          // 0x0805
		"PKCS1WithSHA384",        // 0x0501
		"PSSWithSHA512",          // 0x0806
		"PKCS1WithSHA512",        // 0x0601
	}
	supportedVersions := []string{"1.3", "1.2"}
	keyShareCurves := []string{"X25519"}

	supportedProtocolsALPN := []string{"h2", "http/1.1"}
	supportedProtocolsALPS := []string{"h2"}

	echCandidateCipherSuites := []tlsclient.CandidateCipherSuites{
		{
			KdfId:  "HKDF_SHA256",
			AeadId: "AEAD_AES_128_GCM",
		},
		{
			KdfId:  "HKDF_SHA256",
			AeadId: "AEAD_CHACHA20_POLY1305",
		},
	}

	candidatePayloads := []uint16{128, 160, 192, 224}

	specFunc, err := tlsclient.GetSpecFactoryFromJa3String(ja3String, supportedSignatureAlgorithms, supportedDelegatedCredentialsAlgorithms, supportedVersions, keyShareCurves, supportedProtocolsALPN, supportedProtocolsALPS, echCandidateCipherSuites, candidatePayloads, "zlib")

	if err != nil {
		specFunc = profiles.Chrome_124.GetClientHelloSpec
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

	seed, err := bfutls.NewPRNGSeed()
	if err != nil {
		seed = nil
	}

	return profiles.NewClientProfile(bfutls.ClientHelloID{
		Client:               "Edge",
		Version:              "117",
		RandomExtensionOrder: false,
		Seed:                 seed,
		Weights:              &bfutls.DefaultWeights,
		SpecFactory:          specFunc,
	}, settings, settingsOrder, pseudoHeaderOrder, connectionFlow, nil, nil)
}
