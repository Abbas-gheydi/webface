package web

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	UpStream string
	proxy    = NewProxy(os.Getenv("UPSTREAM"))
)

func NewProxy(targetHost string) *httputil.ReverseProxy {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil
	}
	return httputil.NewSingleHostReverseProxy(url)
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	SetInsecureSkipVerifyPolicy()
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func SetInsecureSkipVerifyPolicy() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: InsecureSkipVerify}
}
