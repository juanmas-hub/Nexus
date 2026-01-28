package http

import "net/http"

func (handler *GatewayHandler) setAuthCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   handler.isProd,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600,
	})
}