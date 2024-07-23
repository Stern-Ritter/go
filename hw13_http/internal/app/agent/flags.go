package agent

import (
	"flag"

	"github.com/Stern-Ritter/go/hw13_http/internal/config/agent"
)

func GetConfig() *agent.Config {
	cfg := &agent.Config{
		LoggerLevel: "INFO",
	}
	parseFlags(cfg)

	return cfg
}

func parseFlags(cfg *agent.Config) {
	flag.StringVar(&cfg.ServerURL, "u", "http://localhost:8080", "server url to send requests")
	flag.StringVar(&cfg.ResourceEndpoint, "e", "users", "resource endpoint")
	flag.Parse()
}
