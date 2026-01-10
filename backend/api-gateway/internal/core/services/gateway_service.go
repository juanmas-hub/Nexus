package services

import (
    "context"
    "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
    "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/ports"
)

type GatewayService struct {
    authClient ports.AuthClient
}

func NewGatewayService(ac ports.AuthClient) *GatewayService {
    return &GatewayService{
        authClient: ac,
    }
}

func (service *GatewayService) Login(ctx context.Context, request domain.LoginRequest) (*domain.LoginResponse, error) {
    response, err := service.authClient.Authenticate(ctx, request)
    if err != nil {
        return nil, err
    }

    return response, nil
}

/*func (s *GatewayService) Register(w http.ResponseWriter, r *http.Request) {
	s.authProxy.Forward(w, r, "/auth/register")
}

// CATALOG
func (s *GatewayService) GetEvents(w http.ResponseWriter, r *http.Request) {
    s.catalogProxy.Forward(w, r, "/events")
}*/