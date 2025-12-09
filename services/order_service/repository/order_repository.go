package repository

import (
	"database/sql"
	"order_service/domain"
	"time"
)

func (r *Repository) Create(order *domain.Order) error {
	query := `
		INSERT INTO orders (user_id, product_id, quantity, total_price, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at`
	
	now := time.Now()
	return r.db.QueryRow(query,
		order.UserID,
		order.ProductID,
		order.Quantity,
		order.TotalPrice,
		order.Status,
		now,
		now,
	).Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt)
}

func (r *Repository) FindByID(id int64) (*domain.Order, error) {
	query := `
		SELECT id, user_id, product_id, quantity, total_price, status, created_at, updated_at
		FROM orders
		WHERE id = $1`
	
	order := &domain.Order{}
	err := r.db.QueryRow(query, id).Scan(
		&order.ID,
		&order.UserID,
		&order.ProductID,
		&order.Quantity,
		&order.TotalPrice,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	
	return order, nil
}

func (r *Repository) FindByUserID(userID int64) ([]*domain.Order, error) {
	query := `
		SELECT id, user_id, product_id, quantity, total_price, status, created_at, updated_at
		FROM orders
		WHERE user_id = $1
		ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var orders []*domain.Order
	for rows.Next() {
		order := &domain.Order{}
		err := rows.Scan(
			&order.ID,
			&order.UserID,
			&order.ProductID,
			&order.Quantity,
			&order.TotalPrice,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	
	return orders, nil
}

func (r *Repository) Update(order *domain.Order) error {
	query := `
		UPDATE orders
		SET user_id = $1, product_id = $2, quantity = $3, total_price = $4, status = $5, updated_at = $6
		WHERE id = $7`
	
	order.UpdatedAt = time.Now()
	_, err := r.db.Exec(query,
		order.UserID,
		order.ProductID,
		order.Quantity,
		order.TotalPrice,
		order.Status,
		order.UpdatedAt,
		order.ID,
	)
	
	return err
}

func (r *Repository) Delete(id int64) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}