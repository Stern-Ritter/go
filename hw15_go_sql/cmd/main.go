package main

import (
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/app"
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := app.GetConfig()
	lg := logger.GetLogger(cfg.LoggerLevel)

	err := app.Run(cfg, lg)
	if err != nil {
		lg.WithFields(logrus.Fields{"error": err}).
			Error("Error starting server")
	}
}
