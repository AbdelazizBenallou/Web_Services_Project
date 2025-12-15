package usecase

import (
	"errors"
	"order_service/domain"
	"order_service/repository"
)

type OrderUseCase interface {
	CreateOrder(userID int64, items []domain.OrderItem) (*domain.Order, error)
}

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

func NewOrderUseCase(orderRepo repository.OrderRepository) OrderUseCase {
	return &orderUseCase{orderRepo: orderRepo}
}

func (uc *orderUseCase) CreateOrder(
	userID int64,
	items []domain.OrderItem,
) (*domain.Order, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}
	if len(items) == 0 {
		return nil, errors.New("order must have items")
	}

	order := &domain.Order{
		UserID: userID,
		Items:  items,
	}

	err := uc.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

