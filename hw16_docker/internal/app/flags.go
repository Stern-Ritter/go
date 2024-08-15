package app

import (
	"flag"

	"github.com/Stern-Ritter/go/hw16_docker/internal/config"
	"github.com/caarlos0/env"
)

func GetConfig() (*config.Config, error) {
	cfg := &config.Config{
		LoggerLevel: "INFO",
	}
	parseFlags(cfg)

	err := env.Parse(cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func parseFlags(cfg *config.Config) {
	flag.StringVar(&cfg.Host, "h", "", "server host")
	flag.IntVar(&cfg.Port, "p", 8080, "server port")
	flag.StringVar(&cfg.DatabaseDSN, "d", "", "database dsn")
	flag.Parse()
}
