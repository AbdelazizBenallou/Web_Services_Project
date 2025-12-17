package usecase

import "product_service/domain"

type StockUseCase interface {
	Add(productID int64, qty int) error
	GetByProductID(productID int64) (*domain.Stock, error)
	ReserveForOrder(items []domain.OrderItem) error

}


