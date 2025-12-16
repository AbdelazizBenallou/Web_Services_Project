package repository

import (
	"database/sql"
	"product_service/domain"
)

type productPostgres struct {
	db *sql.DB
}

func NewProductPostgres(db *sql.DB) ProductRepository {
	return &productPostgres{db: db}
}

func (r *productPostgres) Create(product *domain.Product) error {
	return r.db.QueryRow(
		`INSERT INTO products (name, category_id, price)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		product.Name,
		product.CategoryID,
		product.Price,
	).Scan(&product.ID)
}

func (r *productPostgres) GetByID(id int64) (*domain.Product, error) {
	row := r.db.QueryRow(
		`SELECT id, name, category_id, price
		 FROM products WHERE id = $1`,
		id,
	)

	var p domain.Product
	if err := row.Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productPostgres) GetAll() ([]*domain.Product, error) {
	rows, err := r.db.Query(
		`SELECT id, name, category_id, price
		 FROM products ORDER BY id DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}

func (r *productPostgres) GetByCategory(categoryID int64) ([]*domain.Product, error) {
	rows, err := r.db.Query(
		`SELECT id, name, category_id, price
		 FROM products
		 WHERE category_id = $1
		 ORDER BY id DESC`,
		categoryID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.CategoryID, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}

