package repository

import "product_service/domain"

type ProductRepository interface {
	Create(product *domain.Product) error
	GetByID(id int64) (*domain.Product, error)
	GetAll() ([]*domain.Product, error)
	GetByCategory(categoryID int64) ([]*domain.Product, error)
}

