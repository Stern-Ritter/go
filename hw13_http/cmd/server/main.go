package main

import (
	"github.com/Stern-Ritter/go/hw13_http/internal/app/server"
	"github.com/Stern-Ritter/go/hw13_http/internal/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := server.GetConfig()
	lg := logger.GetLogger(cfg.LoggerLevel)

	err := server.Run(cfg, lg)
	if err != nil {
		lg.WithFields(logrus.Fields{"error": err}).
			Error("Error starting server")
	}
}
