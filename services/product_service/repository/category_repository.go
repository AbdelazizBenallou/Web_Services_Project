package repository

import "product_service/domain"

type CategoryRepository interface {
	Create(category *domain.Category) error
	GetAll() ([]*domain.Category, error)
	GetByID(id int64) (*domain.Category, error)

	ExistsByID(id int64) (bool, error)
	ExistsByName(name string) (bool, error)
}

