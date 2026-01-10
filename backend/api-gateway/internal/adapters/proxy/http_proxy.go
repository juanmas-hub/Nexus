package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HTTPProxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewHTTPProxy(targetURL string) (*HTTPProxy, error) {
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}
	return &HTTPProxy{
		target: parsedURL,
		proxy:  httputil.NewSingleHostReverseProxy(parsedURL),
	}, nil
}

func (p *HTTPProxy) Forward(w http.ResponseWriter, r *http.Request, targetPath string) {
	r.URL.Path = targetPath
	r.Host = p.target.Host
	p.proxy.ServeHTTP(w, r)
}