package ports

import (
    "context"
    "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
)

type CatalogClient interface {
    GetEvents(ctx context.Context) ([]domain.Event, error)
}