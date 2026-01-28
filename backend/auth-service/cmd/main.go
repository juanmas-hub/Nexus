package main

import (
	"nexus/auth-service/internal/infra"
	"nexus/auth-service/internal/config"
)

func main() {
    cfg := config.Load()

    authRouter := setupDependencyInjection(cfg)

    infra.StartServer(cfg.Port, authRouter)
}