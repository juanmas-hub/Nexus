package infra

import(
	"log"
	"net/http"
)

func StartServer(port string, router http.Handler) {
    serverAddress := ":" + port
    log.Printf("[AUTH-SERVICE START] Escuchando en %s", serverAddress)

    if err := http.ListenAndServe(serverAddress, router); err != nil {
        log.Fatalf("[CRITICAL ERROR] El servidor falló: %v", err)
    }
}