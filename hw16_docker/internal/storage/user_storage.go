package storage

import (
	"context"
	"database/sql"

	"github.com/Stern-Ritter/go/hw16_docker/internal/model"
	"github.com/sirupsen/logrus"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user model.User) (int64, error)
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, userID int64) error
	GetUserByID(ctx context.Context, userID int64) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

type DBUserStorage struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewUserStorage(db *sql.DB, logger *logrus.Logger) UserStorage {
	return &DBUserStorage{
		db:     db,
		logger: logger,
	}
}

func (s *DBUserStorage) CreateUser(ctx context.Context, user model.User) (int64, error) {
	var userID int64
	err := s.db.QueryRowContext(ctx, `
		INSERT INTO shop.users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`, user.Name, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return -1, err
	}

	return userID, nil
}

func (s *DBUserStorage) UpdateUser(ctx context.Context, user model.User) error {
	_, err := s.db.ExecContext(ctx, `
	UPDATE shop.users
	SET name = $1, email = $2, password = $3 
	WHERE id = $4
`, user.Name, user.Email, user.Password, user.ID)

	return err
}

func (s *DBUserStorage) DeleteUser(ctx context.Context, userID int64) error {
	_, err := s.db.ExecContext(ctx, `
	DELETE FROM shop.users WHERE id = $1
`, userID)

	return err
}

func (s *DBUserStorage) GetUserByID(ctx context.Context, userID int64) (model.User, error) {
	row := s.db.QueryRowContext(ctx, `
	SELECT id, name, email, password
	FROM shop.users
	WHERE id = $1
`, userID)

	user := model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s *DBUserStorage) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	row := s.db.QueryRowContext(ctx, `
	SELECT id, name, email, password
	FROM shop.users
	WHERE email = $1
`, email)

	user := model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
