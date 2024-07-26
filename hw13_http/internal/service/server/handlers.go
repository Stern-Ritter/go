package server

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	er "github.com/Stern-Ritter/go/hw13_http/internal/errors"
	"github.com/Stern-Ritter/go/hw13_http/internal/model"
	"github.com/go-chi/chi/v5"
)

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	createUserDto := model.CreateUserDto{}
	err = json.Unmarshal(data, &createUserDto)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	userDto, err := s.userService.CreateUser(createUserDto)
	if err != nil {
		http.Error(w, "Unexpected internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	body, err := json.Marshal(userDto)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Error writing response body", http.StatusInternalServerError)
	}
}

func (s *Server) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	pathID := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(pathID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	userDto, err := s.userService.GetUser(id)
	var notFoundErr *er.NotFoundError
	if err != nil {
		if errors.As(err, &notFoundErr) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Unexpected internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, err := json.Marshal(userDto)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Error writing response body", http.StatusInternalServerError)
	}
}
