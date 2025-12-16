package usecase

import (
	"errors"
	"product_service/domain"
	"product_service/repository"
)

type productUseCase struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
	stockRepo    repository.StockRepository
}

func NewProductUseCase(
	productRepo repository.ProductRepository,
	categoryRepo repository.CategoryRepository,
	stockRepo repository.StockRepository,
) ProductUseCase {
	return &productUseCase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		stockRepo:    stockRepo,
	}
}

func (uc *productUseCase) Create(
	name string,
	categoryID int64,
	price float64,
	initialStock int,
) (*domain.Product, error) {

	if name == "" {
		return nil, errors.New("product name is required")
	}
	if price <= 0 {
		return nil, errors.New("invalid price")
	}
	if initialStock < 0 {
		return nil, errors.New("invalid stock")
	}

	// ✅ category must exist
	exists, err := uc.categoryRepo.ExistsByID(categoryID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("category does not exist")
	}

	product := &domain.Product{
		Name:       name,
		CategoryID: categoryID,
		Price:      price,
	}

	if err := uc.productRepo.Create(product); err != nil {
		return nil, err
	}

	stock := &domain.Stock{
		ProductID: product.ID,
		Quantity:  initialStock,
	}
	if err := uc.stockRepo.Create(stock); err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *productUseCase) GetByID(id int64) (*domain.Product, error) {
	if id <= 0 {
		return nil, errors.New("invalid product id")
	}
	return uc.productRepo.GetByID(id)
}

func (uc *productUseCase) GetAll() ([]*domain.Product, error) {
	return uc.productRepo.GetAll()
}

func (uc *productUseCase) GetByCategory(categoryID int64) ([]*domain.Product, error) {
	if categoryID <= 0 {
		return nil, errors.New("invalid category id")
	}
	return uc.productRepo.GetByCategory(categoryID)
}

