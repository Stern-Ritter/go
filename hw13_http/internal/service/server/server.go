package server

import (
	"github.com/Stern-Ritter/go/hw13_http/internal/config/server"
	"github.com/sirupsen/logrus"
)

type Server struct {
	userService UserService
	config      *server.Config
	Logger      *logrus.Logger
}

func NewServer(userService UserService, cfg *server.Config, logger *logrus.Logger) *Server {
	return &Server{
		userService: userService,
		config:      cfg,
		Logger:      logger,
	}
}
