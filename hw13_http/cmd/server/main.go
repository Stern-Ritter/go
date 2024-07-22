package main

import (
	"github.com/Stern-Ritter/go/hw13_http/internal/app/server"
	"github.com/Stern-Ritter/go/hw13_http/internal/logger"
)

func main() {
	cfg := server.GetConfig()
	lg := logger.GetLogger(cfg.LoggerLevel)

	err := server.Run(cfg, lg)
	if err != nil {
		lg.Error("Error starting server", "error", err)
	}
}
