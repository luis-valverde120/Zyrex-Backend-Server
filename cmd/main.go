package main

import (
	"github.com/luis-valverde120/Zyrex-Backend-Server/internal/infrastructure/config"
	"github.com/luis-valverde120/Zyrex-Backend-Server/internal/infrastructure/server"
)

func main() {
	cfg := config.Load()

	router := server.NewServer()
	server.StartServer(router, cfg.Port)
}
