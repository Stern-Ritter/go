package agent

import (
	"net/http"

	"github.com/Stern-Ritter/go/hw13_http/internal/config/agent"
	"github.com/sirupsen/logrus"
)

type Agent struct {
	client *http.Client
	config *agent.Config
	Logger *logrus.Logger
}

func NewAgent(client *http.Client, cfg *agent.Config, lg *logrus.Logger) *Agent {
	return &Agent{
		client: client,
		config: cfg,
		Logger: lg,
	}
}
