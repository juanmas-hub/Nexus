package clients

import (
    "context"
    //"net/http"
	"time"

    "github.com/hashicorp/go-retryablehttp"
    "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
)

type HTTPAuthClient struct {
    baseURL    string
    httpClient *retryablehttp.Client
}

func NewHTTPAuthClient(url string, timeout time.Duration) *HTTPAuthClient {
    retryClient := retryablehttp.NewClient()

    retryClient.RetryMax = 10                   
    retryClient.RetryWaitMin = 1 * time.Second 
    retryClient.RetryWaitMax = 15 * time.Second
    retryClient.Logger = nil
    
    retryClient.HTTPClient.Timeout = timeout

    return &HTTPAuthClient{
        baseURL:    url,
        httpClient: retryClient,
    }
}

func (c *HTTPAuthClient) Authenticate(ctx context.Context, request domain.LoginRequest) (*domain.LoginResponse, error) {
    url := c.baseURL + "/auth/login"
    
    return doRequest[domain.LoginResponse](ctx, c.httpClient, "POST", url, request)
}

func (c *HTTPAuthClient) Register(ctx context.Context, request domain.RegisterRequest) (*domain.RegisterResponse, error) {
    url := c.baseURL + "/auth/register"

    return doRequest[domain.RegisterResponse](ctx, c.httpClient, "POST", url, request)
}