package services

import (
	"net/http"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/ports"
)

type GatewayService struct {
	authProxy ports.ProxyProvider
}

func NewGatewayService(auth ports.ProxyProvider) *GatewayService {
	return &GatewayService{
		authProxy: auth,
	}
}

func (s *GatewayService) Login(w http.ResponseWriter, r *http.Request) {
	s.authProxy.Forward(w, r, "/auth/login")
}

func (s *GatewayService) Register(w http.ResponseWriter, r *http.Request) {
	s.authProxy.Forward(w, r, "/auth/register")
}