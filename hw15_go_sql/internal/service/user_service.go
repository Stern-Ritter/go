package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	er "github.com/Stern-Ritter/go/hw15_go_sql/internal/errors"
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/model"
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/storage"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(ctx context.Context, userDto model.CreateUserDto) (model.UserDto, error)
	UpdateUser(ctx context.Context, userDto model.UpdateUserDto, userID int64) (model.UserDto, error)
	DeleteUser(ctx context.Context, userID int64) (model.UserDto, error)
	GetUserByEmail(ctx context.Context, email string) (model.UserDto, error)
	GetUser(ctx context.Context, userID int64) (model.User, error)
}

type UserServiceImpl struct {
	userStorage storage.UserStorage
	Logger      *logrus.Logger
}

func NewUserService(userStorage storage.UserStorage, logger *logrus.Logger) UserService {
	return &UserServiceImpl{
		userStorage: userStorage,
		Logger:      logger,
	}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, createUserDto model.CreateUserDto) (model.UserDto, error) {
	user := model.CreateUserDtoToUser(createUserDto)
	userID, err := u.userStorage.CreateUser(ctx, user)
	if err != nil {
		u.Logger.WithError(err).Errorf("Failed to create user: %v", createUserDto)
		return model.UserDto{}, err
	}

	user.ID = userID
	userDto := model.UserToUserDto(user)

	u.Logger.Infof("User created successfully: %v", userDto)
	return userDto, nil
}

//nolint:dupl
func (u *UserServiceImpl) UpdateUser(ctx context.Context, updateUserDto model.UpdateUserDto,
	userID int64,
) (model.UserDto, error) {
	user, err := u.GetUser(ctx, userID)
	if err != nil {
		return model.UserDto{}, err
	}

	updatedUser := model.UpdateUserDtoToUser(updateUserDto, user)
	updatedUser.ID = userID

	err = u.userStorage.UpdateUser(ctx, updatedUser)
	if err != nil {
		u.Logger.WithError(err).Errorf("Failed to update user: %v", updateUserDto)
		return model.UserDto{}, err
	}

	savedUser, err := u.userStorage.GetUserByID(ctx, userID)
	if err != nil {
		u.Logger.WithError(err).Errorf("Failed to get user by ID: %d", userID)
		return model.UserDto{}, err
	}
	userDto := model.UserToUserDto(savedUser)

	u.Logger.Infof("User updated successfully: %v", userDto)
	return userDto, nil
}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, userID int64) (model.UserDto, error) {
	user, err := u.GetUser(ctx, userID)
	if err != nil {
		return model.UserDto{}, err
	}

	err = u.userStorage.DeleteUser(ctx, userID)
	if err != nil {
		u.Logger.WithError(err).Errorf("Failed to delete user: %v", user)
		return model.UserDto{}, err
	}
	userDto := model.UserToUserDto(user)

	u.Logger.Infof("User deleted successfully: %v", userDto)
	return userDto, nil
}

func (u *UserServiceImpl) GetUserByEmail(ctx context.Context, email string) (model.UserDto, error) {
	user, err := u.userStorage.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			u.Logger.WithError(err).Infof("User with email: %s doesn't exist", email)
			return model.UserDto{}, er.NewNotFoundError(fmt.Sprintf("User with email: %s doesn`t exist", email), err)
		}
		u.Logger.WithError(err).Errorf("Failed to get user by email: %s", email)
		return model.UserDto{}, err
	}
	userDto := model.UserToUserDto(user)

	u.Logger.Infof("User found successfully: %v", userDto)
	return userDto, nil
}

func (u *UserServiceImpl) GetUser(ctx context.Context, userID int64) (model.User, error) {
	user, err := u.userStorage.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			u.Logger.WithError(err).Infof("User with ID: %d doesn't exist", userID)
			return model.User{}, er.NewNotFoundError(fmt.Sprintf("User with ID: %d doesn`t exist", userID), err)
		}
		u.Logger.WithError(err).Errorf("Failed to get user by ID: %d", userID)
		return model.User{}, err
	}

	return user, nil
}
