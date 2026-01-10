package ports

import (
    "context"
    "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
)

type AuthClient interface {
    Authenticate(ctx context.Context, request domain.LoginRequest) (*domain.LoginResponse, error)
}