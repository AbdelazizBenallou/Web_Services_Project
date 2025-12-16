package repository

import "product_service/domain"

type StockRepository interface {
	Create(stock *domain.Stock) error
	GetByProductID(productID int64) (*domain.Stock, error)
	Update(stock *domain.Stock) error
}

