package ports

import "net/http"

type ProxyProvider interface {
	Forward(w http.ResponseWriter, r *http.Request, targetPath string)
}