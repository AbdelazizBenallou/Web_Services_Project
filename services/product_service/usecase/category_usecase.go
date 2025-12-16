package usecase

import "product_service/domain"

type CategoryUseCase interface {
	Create(name string) (*domain.Category, error)
	GetAll() ([]*domain.Category, error)
	GetByID(id int64) (*domain.Category, error)
}

