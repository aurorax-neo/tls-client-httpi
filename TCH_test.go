package tls_client_httpi

import (
	"github.com/aurorax-neo/tls_client_httpi/TCHUtil"
	"testing"
)

func TestGetReq(t *testing.T) {
	c := DefaultClient()
	response, err := c.Request("GET", "https://www.baidu.com", nil, nil, nil)
	if err != nil {
		return
	}
	TCHUtil.OutResponse(response)
}
