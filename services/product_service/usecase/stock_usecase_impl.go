package usecase

import (
	"errors"

	"product_service/domain"
	"product_service/repository"
)

type stockUseCase struct {
	repo repository.StockRepository
}

func NewStockUseCase(repo repository.StockRepository) StockUseCase {
	return &stockUseCase{repo: repo}
}

func (uc *stockUseCase) Add(productID int64, qty int) error {
	if productID <= 0 {
		return errors.New("invalid product id")
	}
	if qty <= 0 {
		return errors.New("invalid quantity")
	}

	stock, err := uc.repo.GetByProductID(productID)
	if err != nil {
		return err
	}

	stock.Quantity += qty
	return uc.repo.Update(stock)
}

func (uc *stockUseCase) GetByProductID(productID int64) (*domain.Stock, error) {
	if productID <= 0 {
		return nil, errors.New("invalid product id")
	}
	return uc.repo.GetByProductID(productID)
}

func (uc *stockUseCase) ReserveForOrder(items []domain.OrderItem) error {
	for _, item := range items {
		stock, err := uc.repo.GetByProductID(item.ProductID)
		if err != nil {
			return err
		}

		if err := stock.Reserve(item.Quantity); err != nil {
			return err
		}

		if err := uc.repo.Update(stock); err != nil {
			return err
		}
	}
	return nil
}

