package main

import (
	"github.com/Stern-Ritter/go/hw13_http/internal/app/agent"
	"github.com/Stern-Ritter/go/hw13_http/internal/logger"
)

func main() {
	cfg := agent.GetConfig()
	lg := logger.GetLogger(cfg.LoggerLevel)

	err := agent.Run(cfg, lg)
	if err != nil {
		lg.Error("Error starting agent", "error", err)
	}
}
