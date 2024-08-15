package service

import (
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/config"
	"github.com/sirupsen/logrus"
)

type Server struct {
	userService    UserService
	productService ProductService
	orderService   OrderService
	config         *config.Config
	Logger         *logrus.Logger
}

func NewServer(userService UserService, productService ProductService, orderService OrderService, cfg *config.Config,
	logger *logrus.Logger,
) *Server {
	return &Server{
		userService:    userService,
		productService: productService,
		orderService:   orderService,
		config:         cfg,
		Logger:         logger,
	}
}
