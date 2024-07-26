package server

import (
	"github.com/Stern-Ritter/go/hw13_http/internal/model"
	storage "github.com/Stern-Ritter/go/hw13_http/internal/storage/server"
)

type UserService interface {
	CreateUser(model.CreateUserDto) (model.UserDto, error)
	GetUser(uint64) (model.UserDto, error)
}

type UserServiceImpl struct {
	storage storage.UserStorage
}

func NewUserService(storage storage.UserStorage) UserService {
	return &UserServiceImpl{
		storage: storage,
	}
}

func (s *UserServiceImpl) CreateUser(dto model.CreateUserDto) (model.UserDto, error) {
	user := model.CreateUserDtoToUser(dto)
	user, err := s.storage.CreateUser(user)
	if err != nil {
		return model.UserDto{}, err
	}
	userDto := model.UserToUserDto(user)
	return userDto, nil
}

func (s *UserServiceImpl) GetUser(id uint64) (model.UserDto, error) {
	user, err := s.storage.GetUser(id)
	if err != nil {
		return model.UserDto{}, err
	}
	userDto := model.UserToUserDto(user)
	return userDto, nil
}
