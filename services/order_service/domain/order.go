package domain

import "time"

type OrderStatus string

const (
	StatusPendingInventory OrderStatus = "PENDING_INVENTORY"
	StatusConfirmed        OrderStatus = "CONFIRMED"
	StatusCancelled        OrderStatus = "CANCELLED"
)

type Order struct {
	ID        int64        `json:"id"`
	UserID    int64        `json:"user_id"`
	Status    OrderStatus `json:"status"`
	Items     []OrderItem `json:"items"`
	CreatedAt time.Time   `json:"created_at"`
}

