package agent

import (
	"github.com/Stern-Ritter/go/hw13_http/internal/config/agent"
	"github.com/spf13/pflag"
)

func GetConfig() *agent.Config {
	cfg := &agent.Config{
		LoggerLevel: "INFO",
	}
	parseFlags(cfg)

	return cfg
}

func parseFlags(cfg *agent.Config) {
	pflag.StringVarP(&cfg.ServerURL, "url", "u", "http://localhost:8080", "server url to send requests")
	pflag.StringVarP(&cfg.ResourceEndpoint, "endpoint", "e", "users", "resource endpoint")
	pflag.Parse()
}
