package repository

import (
	"database/sql"
	"product_service/domain"
)

type categoryPostgres struct {
	db *sql.DB
}

func NewCategoryPostgres(db *sql.DB) CategoryRepository {
	return &categoryPostgres{db: db}
}

func (r *categoryPostgres) Create(category *domain.Category) error {
	return r.db.QueryRow(
		`INSERT INTO categories (name)
		 VALUES ($1)
		 RETURNING id`,
		category.Name,
	).Scan(&category.ID)
}

func (r *categoryPostgres) GetAll() ([]*domain.Category, error) {
	rows, err := r.db.Query(
		`SELECT id, name FROM categories ORDER BY name`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.Category
	for rows.Next() {
		var c domain.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *categoryPostgres) GetByID(id int64) (*domain.Category, error) {
	row := r.db.QueryRow(
		`SELECT id, name FROM categories WHERE id = $1`,
		id,
	)

	var c domain.Category
	if err := row.Scan(&c.ID, &c.Name); err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *categoryPostgres) ExistsByID(id int64) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		`SELECT EXISTS (SELECT 1 FROM categories WHERE id = $1)`,
		id,
	).Scan(&exists)
	return exists, err
}

func (r *categoryPostgres) ExistsByName(name string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		`SELECT EXISTS (SELECT 1 FROM categories WHERE name = $1)`,
		name,
	).Scan(&exists)
	return exists, err
}

