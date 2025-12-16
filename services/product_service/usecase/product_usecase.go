package usecase

import "product_service/domain"

type ProductUseCase interface {
	Create(
		name string,
		categoryID int64,
		price float64,
		initialStock int,
	) (*domain.Product, error)

	GetByID(id int64) (*domain.Product, error)
	GetAll() ([]*domain.Product, error)
	GetByCategory(categoryID int64) ([]*domain.Product, error)
}

