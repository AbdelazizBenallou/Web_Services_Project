package usecase

import (
	"errors"
	"fmt"
	"product_service/domain"
	"product_service/repository"
)

// تعريف الأخطاء الموحدة
var (
	ErrProductNotFound   = errors.New("product not found")
	ErrInvalidProductID  = errors.New("invalid product id")
	ErrInvalidQuantity   = errors.New("quantity must be greater than 0")
)

// ProductUsecase encapsulates product-related business logic.
type ProductUsecase struct {
	repo repository.ProductRepository
}

// NewProductUsecase creates a new ProductUsecase instance.
func NewProductUsecase(repo repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

// GetByID retrieves a product by ID.
func (p *ProductUsecase) GetByID(id string) (*domain.Product, error) {
	if id == "" {
		return nil, ErrInvalidProductID
	}

	product, err := p.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, ErrProductNotFound
	}

	return product, nil
}

// GetAll retrieves all products.
func (p *ProductUsecase) GetAll() ([]*domain.Product, error) {
	return p.repo.GetAll()
}

// Create creates a new product.
func (p *ProductUsecase) Create(product *domain.Product) error {
	if product.Name == "" || product.Price <= 0 {
		return errors.New("name and price are required and must be greater than 0")
	}
	return p.repo.Create(product)
}

// Update updates an existing product.
func (p *ProductUsecase) Update(product *domain.Product) error {
	if product.ID == "" {
		return ErrInvalidProductID
	}

	existing, err := p.repo.GetByID(product.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrProductNotFound
	}

	return p.repo.Update(product)
}

// Delete deletes a product by ID.
func (p *ProductUsecase) Delete(id string) error {
	if id == "" {
		return ErrInvalidProductID
	}

	existing, err := p.repo.GetByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrProductNotFound
	}

	return p.repo.Delete(id)
}

// UpdateStock updates the stock quantity of a product.
func (p *ProductUsecase) UpdateStock(id string, quantity int) error {
	if id == "" {
		return ErrInvalidProductID
	}
	if quantity < 0 {
		return ErrInvalidQuantity
	}

	existing, err := p.repo.GetByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrProductNotFound
	}

	return p.repo.UpdateStock(id, quantity)
}
