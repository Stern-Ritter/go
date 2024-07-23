package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Stern-Ritter/go/hw13_http/internal/config/server"
	service "github.com/Stern-Ritter/go/hw13_http/internal/service/server"
	storage "github.com/Stern-Ritter/go/hw13_http/internal/storage/server"
	"github.com/go-chi/chi/v5"
)

func Run(cfg *server.Config, log *slog.Logger) error {
	userStorage := storage.NewUserStorage()
	userService := service.NewUserService(userStorage)
	srv := service.NewServer(userService, cfg, log)

	r := chi.NewRouter()
	addRoutes(r, srv)

	url := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	server := &http.Server{
		Handler:      r,
		Addr:         url,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Info("Server running", "url", url)
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

func addRoutes(router chi.Router, srv *service.Server) {
	router.Use(srv.LoggerMiddleware)
	router.Post("/users", srv.CreateUserHandler)
	router.Get("/users/{id}", srv.GetUserHandler)
}
