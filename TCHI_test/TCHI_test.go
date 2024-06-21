package TCHI_test

import (
	"fmt"
	"github.com/aurorax-neo/tls_client_httpi/TCHUtil"
	"github.com/aurorax-neo/tls_client_httpi/cycle_tls"
	"github.com/aurorax-neo/tls_client_httpi/tls_client"
	"testing"
)

func TestGetReq(t *testing.T) {
	c := tls_client.DefaultClient()
	response, err := c.Request("GET", "https://tls.browserleaks.com/json", nil, nil, nil)
	if err != nil {
		return
	}
	TCHUtil.OutHttpResponse(response)

	cc := cycle_tls.DefaultClient()
	response, err = cc.Request("GET", "https://tls.browserleaks.com/json", nil, nil, nil)
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

	cc := cycle_tls.DefaultClient()
	cc.SetProxy("http://127.0.0.1:7890")
	response, err = cc.Request("GET", "https://www.ip.cn/api/index?ip&type=0", nil, nil, nil)
	if err != nil {
		return
	}
	fmt.Println("cc")
	TCHUtil.OutHttpResponse(response)
}
