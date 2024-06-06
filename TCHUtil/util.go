package TCHUtil

import (
	"fmt"
	fhttp "github.com/bogdanfinn/fhttp"
	fhttputil "github.com/bogdanfinn/fhttp/httputil"
	"net/http"
	"net/http/httputil"
)

// OutHttpRequest 打印请求.
func OutHttpRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
	} else {
		fmt.Println(string(dump))
	}
}

// OutHttpResponse 打印响应.
func OutHttpResponse(res *http.Response) {
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("Error dumping response:", err)
	} else {
		fmt.Println(string(dump))
	}
}

// OutFHttpRequest 打印请求.
func OutFHttpRequest(req *fhttp.Request) {
	dump, err := fhttputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
	} else {
		fmt.Println(string(dump))
	}
}

// OutFHttpResponse 打印响应.
func OutFHttpResponse(res *fhttp.Response) {
	dump, err := fhttputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("Error dumping response:", err)
	} else {
		fmt.Println(string(dump))
	}
}
