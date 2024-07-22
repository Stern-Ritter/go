package server

import (
	"github.com/Stern-Ritter/go/hw13_http/internal/config/server"
	"github.com/spf13/pflag"
)

func GetConfig() *server.Config {
	cfg := &server.Config{
		LoggerLevel: "INFO",
	}
	parseFlags(cfg)

	return cfg
}

func parseFlags(c *server.Config) {
	pflag.StringVarP(&c.Host, "host", "h", "localhost", "server host")
	pflag.IntVarP(&c.Port, "port", "p", 8080, "server port")
	pflag.Parse()
}
