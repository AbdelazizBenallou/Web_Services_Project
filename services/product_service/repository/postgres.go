package repository

import (
	"database/sql"
	"product_service/domain"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ProductRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetByID(id string) (*domain.Product, error) {
	query := `SELECT id, name, description, price, stock, created_at, updated_at 
	          FROM products WHERE id = $1`
	
	row := r.db.QueryRow(query, id)
	
	var product domain.Product
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	
	if err == sql.ErrNoRows {
		return nil, nil 
	}
	
	if err != nil {
		return nil, err
	}
	
	return &product, nil
}

func (r *postgresRepository) GetAll() ([]*domain.Product, error) {
	query := `SELECT id, name, description, price, stock, created_at, updated_at 
	          FROM products ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	
	return products, nil
}

func (r *postgresRepository) Create(product *domain.Product) error {
	query := `INSERT INTO products (id, name, description, price, stock) 
	          VALUES ($1, $2, $3, $4, $5)`
	
	_, err := r.db.Exec(query,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	)
	
	return err
}

func (r *postgresRepository) Update(product *domain.Product) error {
	query := `UPDATE products 
	          SET name = $1, description = $2, price = $3, stock = $4, updated_at = CURRENT_TIMESTAMP 
	          WHERE id = $5`
	
	_, err := r.db.Exec(query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.ID,
	)
	
	return err
}

func (r *postgresRepository) Delete(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *postgresRepository) UpdateStock(id string, quantity int) error {
	query := `UPDATE products SET stock = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`
	_, err := r.db.Exec(query, quantity, id)
	return err
}
