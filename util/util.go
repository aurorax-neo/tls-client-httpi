package util

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// OutRequest 打印请求.
func OutRequest(req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Error dumping request:", err)
	} else {
		fmt.Println(string(dump))
	}
}

// OutResponse 打印响应.
func OutResponse(res *http.Response) {
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("Error dumping response:", err)
	} else {
		fmt.Println(string(dump))
	}
}
