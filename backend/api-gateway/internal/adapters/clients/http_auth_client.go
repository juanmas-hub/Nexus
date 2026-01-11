package clients

import (
    "context"
    //"log"
    "net/http"
	"time"

    "github.com/juanmas-hub/nexus/backend/api-gateway/internal/core/domain"
)

type HTTPAuthClient struct {
    baseURL    string
    httpClient *http.Client
}

func NewHTTPAuthClient(url string, timeout time.Duration) *HTTPAuthClient {
    return &HTTPAuthClient{
        baseURL: url,
        httpClient: &http.Client{
            Timeout: timeout,
        },
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