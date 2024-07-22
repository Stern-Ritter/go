package server

import (
	"fmt"
	"sync"
	"sync/atomic"

	er "github.com/Stern-Ritter/go/hw13_http/internal/errors"
	"github.com/Stern-Ritter/go/hw13_http/internal/model"
)

type UserStorage interface {
	CreateUser(model.User) (model.User, error)
	GetUser(uint64) (model.User, error)
}

type InMemoryUserStorage struct {
	muUsers   sync.RWMutex
	users     map[uint64]model.User
	currentID uint64
}

func NewUserStorage() UserStorage {
	return &InMemoryUserStorage{
		users: make(map[uint64]model.User),
	}
}

func (s *InMemoryUserStorage) CreateUser(user model.User) (model.User, error) {
	s.muUsers.Lock()
	defer s.muUsers.Unlock()

	id := s.getNextID()
	user.ID = id
	s.users[id] = user

	return user, nil
}

func (s *InMemoryUserStorage) GetUser(id uint64) (model.User, error) {
	s.muUsers.RLock()
	defer s.muUsers.RUnlock()

	user, ok := s.users[id]
	if !ok {
		return model.User{}, er.NewNotFoundError(fmt.Sprintf("User with id: %d not found", id), nil)
	}

	return user, nil
}

func (s *InMemoryUserStorage) getNextID() uint64 {
	return atomic.AddUint64(&s.currentID, 1)
}
