package storage

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Stern-Ritter/go/hw16_docker/internal/model"
	"github.com/sirupsen/logrus"
)

type ProductStorage interface {
	CreateProduct(ctx context.Context, product model.Product) (int64, error)
	UpdateProduct(ctx context.Context, product model.Product) error
	DeleteProduct(ctx context.Context, productID int64) error
	GetProductByID(ctx context.Context, productID int64) (model.Product, error)
	GetProductsByIDs(ctx context.Context, productIDs []int64) ([]model.Product, error)
	GetAllProductsByPrice(ctx context.Context, min *float64, max *float64) ([]model.Product, error)
}

type DBProductStorage struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewProductStorage(db *sql.DB, logger *logrus.Logger) ProductStorage {
	return &DBProductStorage{
		db:     db,
		logger: logger,
	}
}

func (s *DBProductStorage) CreateProduct(ctx context.Context, product model.Product) (int64, error) {
	var productID int64
	err := s.db.QueryRowContext(ctx, `
		INSERT INTO shop.products (name, price)
		VALUES ($1, $2)
		RETURNING id
	`, product.Name, product.Price).Scan(&productID)
	if err != nil {
		return -1, err
	}

	return productID, nil
}

func (s *DBProductStorage) UpdateProduct(ctx context.Context, product model.Product) error {
	_, err := s.db.ExecContext(ctx, `
	UPDATE shop.products
	SET name = $1, price = $2
	WHERE id = $3
`, product.Name, product.Price, product.ID)

	return err
}

func (s *DBProductStorage) DeleteProduct(ctx context.Context, productID int64) error {
	_, err := s.db.ExecContext(ctx, `
	DELETE FROM shop.products
    WHERE id = $1
`, productID)

	return err
}

func (s *DBProductStorage) GetProductByID(ctx context.Context, productID int64) (model.Product, error) {
	row := s.db.QueryRowContext(ctx, `
	SELECT id, name, price
	FROM shop.products
	WHERE id = $1
`, productID)

	product := model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (s *DBProductStorage) GetProductsByIDs(ctx context.Context, productIDs []int64) ([]model.Product, error) {
	if len(productIDs) == 0 {
		return make([]model.Product, 0), nil
	}

	placeholders := make([]string, len(productIDs))
	args := make([]interface{}, len(productIDs))

	for i, id := range productIDs {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	//nolint:gosec
	query := fmt.Sprintf(`
        SELECT id, name, price
        FROM shop.products
        WHERE id IN (%s)
    `, strings.Join(placeholders, ", "))

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		product := model.Product{}
		if err = rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *DBProductStorage) GetAllProductsByPrice(ctx context.Context, min *float64,
	max *float64,
) ([]model.Product, error) {
	query := `
        SELECT id, name, price
        FROM shop.products
    `
	conditions := []string{}
	argNum := 1
	args := []interface{}{}

	if min != nil {
		conditions = append(conditions, fmt.Sprintf("price >= $%d", argNum))
		args = append(args, *min)
		argNum++
	}
	if max != nil {
		conditions = append(conditions, fmt.Sprintf("price <= $%d", argNum))
		args = append(args, *max)
	}
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		product := model.Product{}
		if err = rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
