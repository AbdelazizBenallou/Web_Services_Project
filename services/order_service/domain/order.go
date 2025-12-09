package domain

import (
	"time"
)

type Order struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	ProductID  int64     `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderRepository interface {
	Create(order *Order) error
	FindByID(id int64) (*Order, error)
	FindByUserID(userID int64) ([]*Order, error)
	Update(order *Order) error
	Delete(id int64) error
}

type OrderUsecase interface {
	CreateOrder(order *Order) error
	GetOrderByID(id int64) (*Order, error)
	GetOrdersByUserID(userID int64) ([]*Order, error)
	UpdateOrder(order *Order) error
	DeleteOrder(id int64) error
}