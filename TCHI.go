package tls_client_httpi

import (
	"io"
	"net/http"
)

type TCHI interface {
	Request(method Method, rawURL string, headers Headers, cookies Cookies, body io.Reader) (*http.Response, error)
	SetProxy(rawUrl string) error
	GetProxy() string
	SetFollowRedirect(followRedirect bool)
	GetFollowRedirect() bool
}

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	HEAD    Method = "HEAD"
	PATCH   Method = "PATCH"
	DELETE  Method = "DELETE"
	OPTIONS Method = "OPTIONS"
)

type Headers map[string]string

func (H Headers) Append(headers Headers) Headers {
	for k, v := range headers {
		H[k] = v
	}
	return H
}
func (H Headers) Set(key, value string) {
	H[key] = value
}
func (H Headers) Get(key string) string {
	return H[key]
}

func (H Headers) Del(key string) {
	delete(H, key)
}

type Cookies []*http.Cookie

func (C Cookies) Append(cookie *http.Cookie) Cookies {
	return append(C, cookie)
}

func (C Cookies) Set(cookie *http.Cookie) Cookies {
	for i, c := range C {
		if c.Name == cookie.Name {
			C[i] = cookie
			return C
		}
	}
	return append(C, cookie)
}
func (C Cookies) Get(name string) *http.Cookie {
	for _, cookie := range C {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

func (C Cookies) Del(name string) Cookies {
	for i, cookie := range C {
		if cookie.Name == name {
			return append(C[:i], C[i+1:]...)
		}
	}
	return C
}
