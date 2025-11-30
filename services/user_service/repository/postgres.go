// Package repository implements data access using PostgreSQL.
package repository

import (
	"database/sql"
	"fmt"
	"user_services/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(connStr string) (*PostgresUserRepository, error) {
	config, err := pgx.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}
	db := stdlib.OpenDB(*config)
	return &PostgresUserRepository{db: db}, nil
}

// GetByID implements UserRepository.GetByID
func (r *PostgresUserRepository) GetByID(id int64) (*domain.User, error) {
	row := r.db.QueryRow(
		`SELECT user_id, full_name, email, password_hash, created_at 
		 FROM users 
		 WHERE user_id = $1`, id)

	var u domain.User
	err := row.Scan(&u.ID, &u.FullName, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found with ID %d", id)
		}
		return nil, fmt.Errorf("database query failed: %w", err)
	}
	return &u, nil
}

// GetAll implements UserRepository.GetAll
func (r *PostgresUserRepository) GetAll() ([]*domain.User, error) {
	rows, err := r.db.Query(
		`SELECT user_id, full_name, email, password_hash, created_at 
		 FROM users`)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.ID, &u.FullName, &u.Email, &u.Password, &u.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, &u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return users, nil
}
