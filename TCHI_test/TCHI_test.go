package TCHI_test

import (
	"fmt"
	"github.com/aurorax-neo/tls_client_httpi"
	"github.com/aurorax-neo/tls_client_httpi/TCHUtil"
	"github.com/aurorax-neo/tls_client_httpi/tls_client"
	"github.com/bogdanfinn/tls-client/profiles"
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

	_ = profiles.Chrome_124

	ccc := tls_client.DefaultClient()
	rr, err := ccc.Request("GET", "https://tls.browserleaks.com/json", nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	TCHUtil.OutHttpResponse(rr)

	cg := tls_client.NewClientOptions(30, tls_client.Edge117())
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
