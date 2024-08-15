package app

import (
	"flag"

	"github.com/Stern-Ritter/go/hw15_go_sql/internal/config"
)

func GetConfig() *config.Config {
	cfg := &config.Config{
		LoggerLevel: "INFO",
	}
	parseFlags(cfg)

	return cfg
}

func parseFlags(cfg *config.Config) {
	flag.StringVar(&cfg.Host, "h", "localhost", "server host")
	flag.IntVar(&cfg.Port, "p", 8080, "server port")
	flag.StringVar(&cfg.DatabaseDSN, "d", "", "database dsn")
	flag.Parse()
}
