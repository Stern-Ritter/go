package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	er "github.com/Stern-Ritter/go/hw15_go_sql/internal/errors"
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/model"
	"github.com/Stern-Ritter/go/hw15_go_sql/internal/storage"
	"github.com/sirupsen/logrus"
)

type OrderService interface {
	CreateOrder(ctx context.Context, orderDto model.CreateOrderDto) (model.OrderDto, error)
	DeleteOrder(ctx context.Context, orderID int64) (model.OrderDto, error)
	GetAllOrdersByUserEmail(ctx context.Context, email string) ([]model.OrderDto, error)
	GetOrdersStatisticByUserEmail(ctx context.Context, email string) (model.OrdersStatistic, error)
}

type OrderServiceImpl struct {
	userService    UserService
	productService ProductService
	orderStorage   storage.OrderStorage
	Logger         *logrus.Logger
}

func NewOrderService(userService UserService, productService ProductService,
	orderStorage storage.OrderStorage, logger *logrus.Logger,
) OrderService {
	return &OrderServiceImpl{
		userService:    userService,
		productService: productService,
		orderStorage:   orderStorage,
		Logger:         logger,
	}
}

func (o *OrderServiceImpl) CreateOrder(ctx context.Context,
	createOrderDto model.CreateOrderDto,
) (model.OrderDto, error) {
	user, err := o.userService.GetUser(ctx, createOrderDto.UserID)
	if err != nil {
		return model.OrderDto{}, err
	}

	products, err := o.productService.GetProductsByIDs(ctx, createOrderDto.ProductsID)
	if err != nil {
		return model.OrderDto{}, err
	}

	if len(products) != len(createOrderDto.ProductsID) {
		o.Logger.WithError(err).Infof("Some products with IDs: %v do not exist", createOrderDto.ProductsID)
		return model.OrderDto{}, er.NewNotFoundError(fmt.Sprintf("Some products with IDs: %v do not exist",
			createOrderDto.ProductsID), err)
	}

	var amount float64
	for _, product := range products {
		amount += product.Price
	}

	order := model.Order{
		UserID:   user.ID,
		Date:     time.Now(),
		Products: products,
		Amount:   amount,
	}

	orderID, err := o.orderStorage.CreateOrder(ctx, order)
	if err != nil {
		o.Logger.WithError(err).Errorf("Failed to create order: %v", order)
		return model.OrderDto{}, err
	}

	order.ID = orderID
	orderDto := model.OrderToOrderDto(order)

	o.Logger.Infof("Order created successfully: %v", order)
	return orderDto, nil
}

func (o *OrderServiceImpl) DeleteOrder(ctx context.Context, orderID int64) (model.OrderDto, error) {
	order, err := o.getOrder(ctx, orderID)
	if err != nil {
		return model.OrderDto{}, err
	}

	err = o.orderStorage.DeleteOrder(ctx, orderID)
	if err != nil {
		o.Logger.WithError(err).Errorf("Failed to delete order: %v", order)
		return model.OrderDto{}, err
	}

	orderDto := model.OrderToOrderDto(order)

	o.Logger.Infof("Order deleted successfully: %v", order)
	return orderDto, nil
}

func (o *OrderServiceImpl) GetAllOrdersByUserEmail(ctx context.Context, email string) ([]model.OrderDto, error) {
	orders, err := o.orderStorage.GetAllOrdersByUserEmail(ctx, email)
	if err != nil {
		o.Logger.WithError(err).Errorf("Failed to get all orders by user email: %s", email)
		return nil, err
	}

	ordersDto := model.OrdersToOrdersDto(orders)

	o.Logger.Infof("Orders by user email found successfully: %v", ordersDto)
	return ordersDto, nil
}

func (o *OrderServiceImpl) GetOrdersStatisticByUserEmail(ctx context.Context,
	email string,
) (model.OrdersStatistic, error) {
	ordersStatistic, err := o.orderStorage.GetOrderStatisticByUserEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			o.Logger.WithError(err).Infof("Failed to get orders statistic by user email: %s", email)
			return model.OrdersStatistic{}, er.NewNotFoundError(fmt.Sprintf("User with email: %s doesn`t exist", email), err)
		}
	}

	o.Logger.Infof("Order statistic by user email found successfully: %v", ordersStatistic)
	return ordersStatistic, nil
}

func (o *OrderServiceImpl) getOrder(ctx context.Context, orderID int64) (model.Order, error) {
	order, err := o.orderStorage.GetByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			o.Logger.WithError(err).Infof("Order with ID: %d doesn't exist", orderID)
			return model.Order{}, er.NewNotFoundError(fmt.Sprintf("Order with ID: %d doesn`t exist", orderID), err)
		}
		o.Logger.WithError(err).Errorf("Failed to get order by ID: %v", orderID)
		return model.Order{}, err
	}

	return order, nil
}
