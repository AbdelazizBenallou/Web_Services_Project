package messaging

import "product_service/domain"

type OrderCreatedEvent struct {
	OrderID int64              `json:"order_id"`
	Items   []domain.OrderItem `json:"items"`
}

type InventoryReservedEvent struct {
	OrderID int64              `json:"order_id"`
	Items   []domain.OrderItem `json:"items"`
}

type InventoryFailedEvent struct {
	OrderID int64 `json:"order_id"`
	Reason  string `json:"reason"`
}

