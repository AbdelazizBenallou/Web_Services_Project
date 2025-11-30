package usecase

import (
	"product_service/domain"
	"product_service/repository"
)

type ProductUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (p *ProductUsecase) GetByID(id string) (*domain.Product, error) {
	return p.repo.GetByID(id)
}

func (p *ProductUsecase) GetAll() ([]*domain.Product, error) {
	return p.repo.GetAll()
}

func (p *ProductUsecase) Create(product *domain.Product) error {
	return p.repo.Create(product)
}

func (p *ProductUsecase) Update(product *domain.Product) error {
	return p.repo.Update(product)
}

func (p *ProductUsecase) Delete(id string) error {
	return p.repo.Delete(id)
}

func (p *ProductUsecase) UpdateStock(id string, quantity int) error {
	return p.repo.UpdateStock(id, quantity)
}
