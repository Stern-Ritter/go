package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	er "github.com/Stern-Ritter/go/hw15_go_sql/internal/errors"
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/model"
	"github.com/go-chi/chi/v5"
)

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	createUserDto := model.CreateUserDto{}
	err = json.Unmarshal(data, &createUserDto)
	if err != nil {
		sendErrorResponse(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	userDto, err := s.userService.CreateUser(r.Context(), createUserDto)
	if err != nil {
		sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusCreated, userDto)
}

//nolint:dupl
func (s *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getPathVariableID(r)
	if err != nil {
		sendErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	updateUserDto := model.UpdateUserDto{}
	err = json.Unmarshal(data, &updateUserDto)
	if err != nil {
		sendErrorResponse(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	userDto, err := s.userService.UpdateUser(r.Context(), updateUserDto, userID)
	if err != nil {
		var notFound er.NotFoundError
		if errors.As(err, &notFound) {
			sendErrorResponse(w, "User not found", http.StatusNotFound)
		} else {
			sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		}
		return
	}

	sendResponse(w, http.StatusOK, userDto)
}

func (s *Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := getPathVariableID(r)
	if err != nil {
		sendErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userDto, err := s.userService.DeleteUser(r.Context(), userID)
	if err != nil {
		var notFound er.NotFoundError
		if errors.As(err, &notFound) {
			sendErrorResponse(w, "User not found", http.StatusNotFound)
		} else {
			sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		}
		return
	}

	sendResponse(w, http.StatusOK, userDto)
}

func (s *Server) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if len(email) == 0 {
		sendErrorResponse(w, "Email parameter is required", http.StatusBadRequest)
		return
	}

	userDto, err := s.userService.GetUserByEmail(r.Context(), email)
	if err != nil {
		var notFound er.NotFoundError
		if errors.As(err, &notFound) {
			sendErrorResponse(w, "User not found", http.StatusNotFound)
		} else {
			sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		}
		return
	}

	sendResponse(w, http.StatusOK, userDto)
}

func (s *Server) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	createProductDto := model.CreateProductDto{}
	err = json.Unmarshal(data, &createProductDto)
	if err != nil {
		sendErrorResponse(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	productDto, err := s.productService.CreateProduct(r.Context(), createProductDto)
	if err != nil {
		sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusCreated, productDto)
}

//nolint:dupl
func (s *Server) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	productID, err := getPathVariableID(r)
	if err != nil {
		sendErrorResponse(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	updateProductDto := model.UpdateProductDto{}
	err = json.Unmarshal(data, &updateProductDto)
	if err != nil {
		sendErrorResponse(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	productDto, err := s.productService.UpdateProduct(r.Context(), updateProductDto, productID)
	if err != nil {
		var notFound er.NotFoundError
		if errors.As(err, &notFound) {
			sendErrorResponse(w, "Product not found", http.StatusNotFound)
		} else {
			sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		}
		return
	}

	sendResponse(w, http.StatusOK, productDto)
}

func (s *Server) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	productID, err := getPathVariableID(r)
	if err != nil {
		sendErrorResponse(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	productDto, err := s.productService.DeleteProduct(r.Context(), productID)
	if err != nil {
		var notFound er.NotFoundError
		if errors.As(err, &notFound) {
			sendErrorResponse(w, "Product not found", http.StatusNotFound)
		} else {
			sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		}
		return
	}

	sendResponse(w, http.StatusOK, productDto)
}

func (s *Server) GetAllProductsByPriceHandler(w http.ResponseWriter, r *http.Request) {
	minPriceStr := r.URL.Query().Get("min")
	maxPriceStr := r.URL.Query().Get("max")

	var min, max *float64
	if minPriceStr != "" {
		minPrice, err := strconv.ParseFloat(minPriceStr, 64)
		if err == nil {
			min = &minPrice
		}
	}
	if maxPriceStr != "" {
		maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
		if err == nil {
			max = &maxPrice
		}
	}

	productsDto, err := s.productService.GetAllProductsByPrice(r.Context(), min, max)
	if err != nil {
		sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusOK, productsDto)
}

func (s *Server) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	createOrderDto := model.CreateOrderDto{}
	err = json.Unmarshal(data, &createOrderDto)
	if err != nil {
		sendErrorResponse(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	orderDto, err := s.orderService.CreateOrder(r.Context(), createOrderDto)
	if err != nil {
		var notFound er.NotFoundError
		if errors.As(err, &notFound) {
			sendErrorResponse(w, err.Error(), http.StatusNotFound)
		} else {
			sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		}
		return
	}

	sendResponse(w, http.StatusCreated, orderDto)
}

func (s *Server) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID, err := getPathVariableID(r)
	if err != nil {
		sendErrorResponse(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	orderDto, err := s.orderService.DeleteOrder(r.Context(), orderID)
	if err != nil {
		var notFound er.NotFoundError
		if errors.As(err, &notFound) {
			sendErrorResponse(w, "Order not found", http.StatusNotFound)
		} else {
			sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		}
		return
	}

	sendResponse(w, http.StatusOK, orderDto)
}

func (s *Server) GetAllOrdersByUserEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if len(email) == 0 {
		sendErrorResponse(w, "Email parameter is required", http.StatusBadRequest)
		return
	}

	ordersDto, err := s.orderService.GetAllOrdersByUserEmail(r.Context(), email)
	if err != nil {
		sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		return
	}

	sendResponse(w, http.StatusOK, ordersDto)
}

func (s *Server) GetOrdersStatisticByUserEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if len(email) == 0 {
		sendErrorResponse(w, "Email parameter is required", http.StatusBadRequest)
		return
	}

	ordersStatistic, err := s.orderService.GetOrdersStatisticByUserEmail(r.Context(), email)
	if err != nil {
		var notFound er.NotFoundError
		if errors.As(err, &notFound) {
			sendErrorResponse(w, "User not found", http.StatusNotFound)
		} else {
			sendErrorResponse(w, "Unexpected internal server error", http.StatusInternalServerError)
		}
		return
	}

	sendResponse(w, http.StatusOK, ordersStatistic)
}

func getPathVariableID(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.ParseInt(idStr, 10, 64)
}

func sendResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	body, err := json.Marshal(data)
	if err != nil {
		sendErrorResponse(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		sendErrorResponse(w, "Error writing response body", http.StatusInternalServerError)
	}
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}
