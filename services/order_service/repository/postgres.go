package repository

import (
	"database/sql"
	"order_service/domain"
)

type Repository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository {
	return &Repository{db: db}
}