package clients

import (
    "context"
    //"log"
    "net/http"
	"time"

    "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
)

type HTTPCatalogClient struct {
    baseURL    string
    httpClient *http.Client
}

func NewHTTPCatalogClient(url string, timeout time.Duration) *HTTPCatalogClient {
    return &HTTPCatalogClient{
        baseURL: url,
        httpClient: &http.Client{
            Timeout: timeout,
        },
    }
}

func (c *HTTPCatalogClient) GetEvents(ctx context.Context) (*domain.GetEventsResponse, error) {
    url := c.baseURL + "/catalog/events"
    
    return doRequest[domain.GetEventsResponse](ctx, c.httpClient, "GET", url, nil)
}