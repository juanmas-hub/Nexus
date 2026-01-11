package clients

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

func doRequest[T any](ctx context.Context, client *http.Client, method, url string, body any) (*T, error) {
    var bodyReader io.Reader
    if body != nil {
        jsonData, err := json.Marshal(body)
        if err != nil {
            return nil, fmt.Errorf("error serializando body: %w", err)
        }
        bodyReader = bytes.NewBuffer(jsonData)
    }

    req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        log.Printf("[GATEWAY ERROR] Fallo de red hacia %s: %v", url, err)
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        errorBody, _ := io.ReadAll(resp.Body)
        log.Printf("[SERVICE ERROR] URL: %s | Status: %d | Body: %s", url, resp.StatusCode, string(errorBody))
        return nil, fmt.Errorf("servicio respondió con status %d", resp.StatusCode)
    }

    var result T
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, fmt.Errorf("error decodificando respuesta: %w", err)
    }

    return &result, nil
}