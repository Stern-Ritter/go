package storage

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Stern-Ritter/go/hw15_go_sql/internal/model"
	"github.com/sirupsen/logrus"
)

type OrderStorage interface {
	CreateOrder(ctx context.Context, order model.Order) (int64, error)
	DeleteOrder(ctx context.Context, orderID int64) error
	GetByID(ctx context.Context, orderID int64) (model.Order, error)
	GetAllOrdersByUserEmail(ctx context.Context, email string) ([]model.Order, error)
	GetOrderStatisticByUserEmail(ctx context.Context, email string) (model.OrdersStatistic, error)
}

type DBOrderStorage struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewOrderStorage(db *sql.DB, logger *logrus.Logger) OrderStorage {
	return &DBOrderStorage{
		db:     db,
		logger: logger,
	}
}

func (s *DBOrderStorage) CreateOrder(ctx context.Context, order model.Order) (int64, error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	var orderID int64
	err = tx.QueryRowContext(ctx, `
		INSERT INTO shop.orders
		(user_id, order_date, total_amount)
		VALUES ($1, $2, $3)
		RETURNING id
	`, order.UserID, order.Date, order.Amount).Scan(&orderID)

	if err != nil {
		return -1, err
	}

	productsID := model.MapProductsToProductsID(order.Products)
	if len(productsID) != 0 {
		placeholders := make([]string, len(productsID))
		args := make([]interface{}, 0, len(productsID)*2)

		for i, productID := range productsID {
			placeholders[i] = fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2)
			args = append(args, orderID, productID)
		}

		//nolint:gosec
		query := fmt.Sprintf(`
        INSERT INTO shop.orders_products (order_id, product_id)
        VALUES %s
		`, strings.Join(placeholders, ", "))

		_, err = tx.ExecContext(ctx, query, args...)
		if err != nil {
			return -1, err
		}
	}

	return orderID, tx.Commit()
}

func (s *DBOrderStorage) DeleteOrder(ctx context.Context, orderID int64) error {
	_, err := s.db.ExecContext(ctx, `
	DELETE FROM shop.orders
	WHERE id = $1
`, orderID)

	return err
}

func (s *DBOrderStorage) GetByID(ctx context.Context, orderID int64) (model.Order, error) {
	row := s.db.QueryRowContext(ctx, `
	SELECT id, user_id, order_date, total_amount
	FROM shop.orders
	WHERE id = $1
`, orderID)

	order := model.Order{}
	err := row.Scan(&order.ID, &order.UserID, &order.Date, &order.Amount)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func (s *DBOrderStorage) GetAllOrdersByUserEmail(ctx context.Context, email string) ([]model.Order, error) {
	rows, err := s.db.QueryContext(ctx, `
	SELECT o.id, u.id, o.order_date, o.total_amount
	FROM shop.users u
	LEFT JOIN shop.orders o
	ON o.user_id = u.id
	WHERE u.email = $1 AND o.id IS NOT NULL
`, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]model.Order, 0)
	for rows.Next() {
		order := model.Order{}
		if err = rows.Scan(&order.ID, &order.UserID, &order.Date, &order.Amount); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *DBOrderStorage) GetOrderStatisticByUserEmail(ctx context.Context,
	email string,
) (model.OrdersStatistic, error) {
	row := s.db.QueryRowContext(ctx, `
	SELECT
	    u.id,
	    u.name,
	    u.email,
       	COALESCE(ot.amount, 0),
       	COALESCE(pt.avg_price,0)
	FROM shop.users u
	LEFT JOIN (
		SELECT user_id, SUM(total_amount) AS amount
		FROM shop.orders
		GROUP BY user_id) ot ON u.id = ot.user_id
	LEFT JOIN (
		SELECT o.user_id, AVG(p.price) AS avg_price
		FROM shop.orders o
		INNER JOIN shop.orders_products op ON o.id = op.order_id
		INNER JOIN shop.products p ON op.product_id = p.id
		GROUP BY o.user_id) pt ON u.id = pt.user_id
	WHERE u.email = $1
`, email)

	orderStatistic := model.OrdersStatistic{}
	err := row.Scan(&orderStatistic.ID, &orderStatistic.Name, &orderStatistic.Email, &orderStatistic.Amount,
		&orderStatistic.AverageProductsPrice)
	if err != nil {
		return model.OrdersStatistic{}, err
	}

	return orderStatistic, nil
}
