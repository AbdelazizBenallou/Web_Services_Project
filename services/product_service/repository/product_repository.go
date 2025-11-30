// repository/product_repository.go
package repository

import "product_service/domain"

type ProductRepository interface {
	GetByID(id string) (*domain.Product, error)
	GetAll() ([]*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id string) error
	UpdateStock(id string, quantity int) error
}
