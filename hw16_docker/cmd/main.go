package main

import (
	"log"

	"github.com/Stern-Ritter/go/hw16_docker/internal/app"
	"github.com/Stern-Ritter/go/hw16_docker/internal/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := app.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %v", err)
	}
	lg := logger.GetLogger(cfg.LoggerLevel)

	err = app.Run(cfg, lg)
	if err != nil {
		lg.WithFields(logrus.Fields{"error": err}).
			Error("Error starting server")
	}
}
