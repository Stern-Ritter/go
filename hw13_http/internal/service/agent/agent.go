package agent

import (
	"log/slog"
	"net/http"

	"github.com/Stern-Ritter/go/hw13_http/internal/config/agent"
)

type Agent struct {
	client *http.Client
	config *agent.Config
	Logger *slog.Logger
}

func NewAgent(client *http.Client, cfg *agent.Config, lg *slog.Logger) *Agent {
	return &Agent{
		client: client,
		config: cfg,
		Logger: lg,
	}
}
