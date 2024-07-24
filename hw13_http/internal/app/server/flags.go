package server

import (
	"flag"

	"github.com/Stern-Ritter/go/hw13_http/internal/config/server"
)

func GetConfig() *server.Config {
	cfg := &server.Config{
		LoggerLevel: "INFO",
	}
	parseFlags(cfg)

	return cfg
}

func parseFlags(c *server.Config) {
	flag.StringVar(&c.Host, "h", "localhost", "server host")
	flag.IntVar(&c.Port, "p", 8080, "server port")
	flag.Parse()
}
