package usecase

import (
	"errors"
	"strings"

	"product_service/domain"
	"product_service/repository"
)

type categoryUseCase struct {
	repo repository.CategoryRepository
}

func NewCategoryUseCase(repo repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{repo: repo}
}

func (uc *categoryUseCase) Create(name string) (*domain.Category, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("category name is required")
	}

	exists, err := uc.repo.ExistsByName(name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("category already exists")
	}

	category := &domain.Category{
		Name: name,
	}

	if err := uc.repo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (uc *categoryUseCase) GetAll() ([]*domain.Category, error) {
	return uc.repo.GetAll()
}

func (uc *categoryUseCase) GetByID(id int64) (*domain.Category, error) {
	if id <= 0 {
		return nil, errors.New("invalid category id")
	}
	return uc.repo.GetByID(id)
}

