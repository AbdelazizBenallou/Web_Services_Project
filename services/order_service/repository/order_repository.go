package repository

import "order_service/domain"

type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id int64) (*domain.Order, error)
	UpdateStatus(orderID int64, status domain.OrderStatus) error

}

