package repository

import (
	"database/sql"
	"order_service/domain"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) OrderRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(order *domain.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.QueryRow(
		`INSERT INTO orders (user_id, status, created_at)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		order.UserID,
		order.Status,
		order.CreatedAt,
	).Scan(&order.ID)

	if err != nil {
		return err
	}

	
	for _, item := range order.Items {
		_, err := tx.Exec(
			`INSERT INTO order_items (order_id, product_id, quantity, price)
			 VALUES ($1, $2, $3, $4)`,
			order.ID,
			item.ProductID,
			item.Quantity,
			item.Price,
		)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}


func (r *postgresRepository) GetByID(id int64) (*domain.Order, error) {
	row := r.db.QueryRow(
		`SELECT id, user_id, status, created_at FROM orders WHERE id=$1`,
		id,
	)

	var o domain.Order
	err := row.Scan(&o.ID, &o.UserID, &o.Status, &o.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (r *postgresRepository) UpdateStatus(id int64, status domain.OrderStatus) error {
	_, err := r.db.Exec(
		`UPDATE orders SET status=$1 WHERE id=$2`,
		status,
		id,
	)
	return err
}

