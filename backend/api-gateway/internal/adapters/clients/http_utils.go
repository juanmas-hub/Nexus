package clients

import (
    //"bytes"
    "context"
    "encoding/json"
    "fmt"
    "github.com/hashicorp/go-retryablehttp"
    //"io"
    //"log"
    //"net/http"
)

func doRequest[T any](ctx context.Context, client *retryablehttp.Client, method, url string, body any) (*T, error) {
    var reqBody interface{}

    if body != nil {
        jsonData, err := json.Marshal(body)
        if err != nil {
            return nil, fmt.Errorf("error serializando body: %w", err)
        }
        reqBody = jsonData 
    }

    req, err := retryablehttp.NewRequestWithContext(ctx, method, url, reqBody)
    if err != nil {
        return nil, fmt.Errorf("error creando request: %w", err)
    }
    
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", "Nexus-Gateway/1.0 (Retryable)")

    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("fallo de comunicación tras reintentos: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return nil, fmt.Errorf("servicio respondió con status %d", resp.StatusCode)
    }

    var result T
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("error decodificando respuesta: %w", err)
    }

    return &result, nil
}