package model

import (
	"time"
)

type Order struct {
	ID       int64
	UserID   int64
	Date     time.Time
	Products []Product
	Amount   float64
}

type CreateOrderDto struct {
	UserID     int64   `json:"userId"`
	ProductsID []int64 `json:"productsId"`
}

type OrderDto struct {
	ID     int64     `json:"id"`
	UserID int64     `json:"userId"`
	Date   time.Time `json:"date"`
	Amount float64   `json:"amount"`
}

type OrdersStatistic struct {
	ID                   int64   `json:"id"`
	Name                 string  `json:"name"`
	Email                string  `json:"email"`
	Amount               float64 `json:"amount"`
	AverageProductsPrice float64 `json:"averageProductsPrice"`
}

func OrderToOrderDto(order Order) OrderDto {
	return OrderDto{
		ID:     order.ID,
		UserID: order.UserID,
		Date:   order.Date,
		Amount: order.Amount,
	}
}

func OrdersToOrdersDto(orders []Order) []OrderDto {
	ordersDto := make([]OrderDto, len(orders))
	for i, order := range orders {
		ordersDto[i] = OrderToOrderDto(order)
	}

	return ordersDto
}

func MapProductsToProductsID(products []Product) []int64 {
	if products == nil {
		return make([]int64, 0)
	}
	productsID := make([]int64, len(products))
	for i, product := range products {
		productsID[i] = product.ID
	}
	return productsID
}
