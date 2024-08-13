package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Stern-Ritter/go/hw15_go_sql/internal/config"
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/service"
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/storage"
	"github.com/Stern-Ritter/go/hw15_go_sql/migrations"
	"github.com/go-chi/chi/v5"
	// Register the pgx driver for database/sql.
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config, log *logrus.Logger) error {
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		log.WithError(err).Fatal("failed open database connection")
	}
	err = migrateDatabase(cfg.DatabaseDSN)
	if err != nil {
		log.WithError(err).Fatal("failed migrate database")
	}

	userStorage := storage.NewUserStorage(db, log)
	productStorage := storage.NewProductStorage(db, log)
	orderStorage := storage.NewOrderStorage(db, log)

	userService := service.NewUserService(userStorage, log)
	productService := service.NewProductService(productStorage, log)
	orderService := service.NewOrderService(userService, productService, orderStorage, log)

	srv := service.NewServer(userService, productService, orderService, cfg, log)

	r := chi.NewRouter()
	addRoutes(srv, r)

	url := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	server := &http.Server{
		Handler:      r,
		Addr:         url,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.WithFields(logrus.Fields{"url": url}).Info("Server running")
	if err = server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

func addRoutes(srv *service.Server, r *chi.Mux) {
	r.Use(srv.LoggerMiddleware)

	r.Post("/users", srv.CreateUserHandler)
	r.Put("/users/{id}", srv.UpdateUserHandler)
	r.Delete("/users/{id}", srv.DeleteUserHandler)
	r.Get("/users", srv.GetUserByEmailHandler)

	r.Post("/products", srv.CreateProductHandler)
	r.Put("/products/{id}", srv.UpdateProductHandler)
	r.Delete("/products/{id}", srv.DeleteProductHandler)
	r.Get("/products", srv.GetAllProductsByPriceHandler)

	r.Post("/orders", srv.CreateOrderHandler)
	r.Delete("/orders/{id}", srv.DeleteOrderHandler)
	r.Get("/orders", srv.GetAllOrdersByUserEmailHandler)
	r.Get("/orders/statistics", srv.GetOrdersStatisticByUserEmailHandler)
}

func migrateDatabase(databaseDsn string) error {
	goose.SetBaseFS(migrations.Migrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose failed to set postgres dialect: %w", err)
	}

	db, err := goose.OpenDBWithDriver("pgx", databaseDsn)
	if err != nil {
		return fmt.Errorf("goose failed to open database connection: %w", err)
	}

	if err = goose.Up(db, "."); err != nil {
		return fmt.Errorf("goose failed to migrate database: %w", err)
	}

	if err = db.Close(); err != nil {
		return fmt.Errorf("goose failed to close database connection: %w", err)
	}

	return nil
}
