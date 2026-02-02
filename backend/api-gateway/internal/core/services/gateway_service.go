package services

import (
    "context"
    "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
    "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/ports"
)

type GatewayService struct {
    authClient ports.AuthClient
    catalogClient ports.CatalogClient
}

func NewGatewayService(ac ports.AuthClient, cc ports.CatalogClient) *GatewayService {
    return &GatewayService{
        authClient: ac,
        catalogClient: cc,
    }
}

func (service *GatewayService) Login(ctx context.Context, request domain.LoginRequest) (*domain.LoginResponse, error) {
    response, err := service.authClient.Authenticate(ctx, request)
    if err != nil {
        return nil, err
    }

    return response, nil
}

func (service *GatewayService) Register(ctx context.Context, request domain.RegisterRequest) (*domain.RegisterResponse, error) {
	response, err := service.authClient.Register(ctx, request)

    if err != nil {
        return nil, err
    }

    return response, nil
}

func (service *GatewayService) GetEvents(ctx context.Context) ([]domain.Event, error){
    response, err := service.catalogClient.GetEvents(ctx)

    if err != nil {
        return nil, err
    }

    return response, nil
}