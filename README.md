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

###### github.com/Danny-Dasilva/CycleTLS/cycletls

```
import (
	"fmt"
	"github.com/aurorax-neo/tls_client_httpi/TCHUtil"
	"github.com/aurorax-neo/tls_client_httpi/cycle_tls"
	"github.com/aurorax-neo/tls_client_httpi/tls_client"
	"testing"
)

func TestGetReq(t *testing.T) {
		cc := cycle_tls.DefaultClient()
	response, err = cc.Request("GET", "https://tls.browserleaks.com/json", nil, nil, nil)
	if err != nil {
		return
	}
	TCHUtil.OutHttpResponse(response)
}
*****************************************************************************************************

HTTP/2.0 200 OK
Content-Length: 692
Access-Control-Allow-Headers: *
Access-Control-Allow-Origin: https://browserleaks.com
Cache-Control: no-store, no-cache, must-revalidate, max-age=0
Content-Type: application/json
Date: Fri, 21 Jun 2024 05:22:04 GMT
Expires: Thu, 19 Nov 1981 08:52:00 GMT
Pragma: no-cache
Server: nginx
X-Content-Type-Options: nosniff

{
  "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
  "ja3_hash": "fe6706a5fa93c16c76a8b6e84fa6e476",
  "ja3_text": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,17513-43-5-35-45-27-10-18-65281-0-23-16-51-11-13-21,29-23-24,0",
  "ja3n_hash": "aa56c057ad164ec4fdcb7a5a283be9fc",
  "ja3n_text": "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-5-10-11-13-16-18-21-23-27-35-43-45-51-17513-65281,29-23-24,0",
  "akamai_hash": "4708d37c97cd9033bbaa2199b0f54c2b",
  "akamai_text": "1:65536;3:1000;4:6291456;5:16384;6:262144|15663105|0|a,m,p,s"
}
```

