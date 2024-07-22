package server

import (
	"log/slog"

	"github.com/Stern-Ritter/go/hw13_http/internal/config/server"
)

type Server struct {
	userService UserService
	config      *server.Config
	Logger      *slog.Logger
}

func NewServer(userService UserService, cfg *server.Config, logger *slog.Logger) *Server {
	return &Server{
		userService: userService,
		config:      cfg,
		Logger:      logger,
	}
}
