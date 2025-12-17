package messaging

import "order_service/domain"

type InventoryReservedEvent struct {
	OrderID int64              `json:"order_id"`
	Items   []domain.OrderItem `json:"items"`
}

type InventoryFailedEvent struct {
	OrderID int64 `json:"order_id"`
	Reason  string `json:"reason"`
}

