package main

import (
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/config"
	"github.com/juanmas-hub/nexus/backend/api-gateway/internal/infra"
)

func main() {
	cfg := config.Load()

	gatewayRouter := setupDependencyInjection(cfg)

	infra.StartServer(cfg.Port, gatewayRouter)
}