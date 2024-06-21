# tls-client-httpi

### 1.安装库

```
go get -u github.com/aurorax-neo/tls_client_httpi
```

### 2.get请求

###### github.com/bogdanfinn/tls-client


```
import (
	"fmt"
	"github.com/aurorax-neo/tls_client_httpi/TCHUtil"
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
}

*****************************************************************************************************

HTTP/2.0 200 OK
Access-Control-Allow-Headers: *
Access-Control-Allow-Origin: https://browserleaks.com
Cache-Control: no-store, no-cache, must-revalidate, max-age=0
Content-Type: application/json
Date: Fri, 21 Jun 2024 05:22:03 GMT
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Pragma: no-cache
Server: nginx
X-Content-Type-Options: nosniff

{
  "user_agent": "Go-http-client/2.0",
  "ja3_hash": "64aff24dbef210f33880d4f62e1493dd",
  "ja3_text": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,27-18-23-17513-16-43-13-11-0-35-10-65037-5-65281-45-51,25497-29-23-24,0",
  "ja3n_hash": "4c9ce26028c11d7544da00d3f7e4f45c",
  "ja3n_text": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-5-10-11-13-16-18-23-27-35-43-45-51-17513-65037-65281,25497-29-23-24,0",
  "akamai_hash": "52d84b11737d980aef856699f885ca86",
  "akamai_text": "1:65536;2:0;4:6291456;6:262144|15663105|0|m,a,s,p"
}
```

