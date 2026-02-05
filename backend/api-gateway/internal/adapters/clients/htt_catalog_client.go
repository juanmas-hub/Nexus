package clients

import (
    "context"
    //"log"
    //"net/http"
	"time"

    "github.com/hashicorp/go-retryablehttp"
    "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
)

type HTTPCatalogClient struct {
    baseURL    string
    httpClient *retryablehttp.Client
}

func NewHTTPCatalogClient(url string, timeout time.Duration) *HTTPCatalogClient {
    retryClient := retryablehttp.NewClient()

    retryClient.RetryMax = 4                   
    retryClient.RetryWaitMin = 1 * time.Second 
    retryClient.RetryWaitMax = 10 * time.Second
    retryClient.Logger = nil
    
    retryClient.HTTPClient.Timeout = timeout

    return &HTTPCatalogClient{
        baseURL:    url,
        httpClient: retryClient,
    }
}

func (c *HTTPCatalogClient) GetEvents(ctx context.Context) ([]domain.Event, error) {
    url := c.baseURL + "/catalog/events"
    
    eventsPtr, err := doRequest[[]domain.Event](ctx, c.httpClient, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    return *eventsPtr, nil
}