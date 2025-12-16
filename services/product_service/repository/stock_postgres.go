package repository

import (
	"database/sql"
	"product_service/domain"
)

type stockPostgres struct {
	db *sql.DB
}

func NewStockPostgres(db *sql.DB) StockRepository {
	return &stockPostgres{db: db}
}

func (r *stockPostgres) Create(stock *domain.Stock) error {
	_, err := r.db.Exec(
		`INSERT INTO stock (product_id, quantity)
		 VALUES ($1, $2)`,
		stock.ProductID,
		stock.Quantity,
	)
	return err
}

func (r *stockPostgres) GetByProductID(productID int64) (*domain.Stock, error) {
	row := r.db.QueryRow(
		`SELECT product_id, quantity
		 FROM stock WHERE product_id = $1`,
		productID,
	)

	var s domain.Stock
	if err := row.Scan(&s.ProductID, &s.Quantity); err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *stockPostgres) Update(stock *domain.Stock) error {
	_, err := r.db.Exec(
		`UPDATE stock
		 SET quantity = $1
		 WHERE product_id = $2`,
		stock.Quantity,
		stock.ProductID,
	)
	return err
}

