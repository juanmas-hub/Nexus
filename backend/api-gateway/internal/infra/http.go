package infra

import(
	"log"
	"net/http"
	
	"github.com/go-chi/chi/v5"
)

func StartServer(port string, router *chi.Mux) {
	serverAddress := ":" + port
	log.Printf("[GATEWAY START] %s", serverAddress)

	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Fatalf("[CRITICAL ERROR] El servidor falló: %v", err)
	}
}