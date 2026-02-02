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

func (c *HTTPCatalogClient) GetEvents(ctx context.Context) ([]domain.Event, error) {
    url := c.baseURL + "/catalog/events"
    
    eventsPtr, err := doRequest[[]domain.Event](ctx, c.httpClient, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    // 3. doRequest devuelve un puntero, así que lo desreferenciamos
    return *eventsPtr, nil
}