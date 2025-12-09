package usecase

import (
	"errors"
	"order_service/domain"
)

type orderUsecase struct {
	orderRepo domain.OrderRepository
}

func NewOrderUsecase(orderRepo domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{orderRepo: orderRepo}
}

func (u *orderUsecase) CreateOrder(order *domain.Order) error {
	if order.UserID <= 0 {
		return errors.New("user_id is required")
	}
	if order.ProductID <= 0 {
		return errors.New("product_id is required")
	}
	if order.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	if order.TotalPrice <= 0 {
		return errors.New("total_price must be greater than 0")
	}
	
	if order.Status == "" {
		order.Status = "pending"
	}
	
	return u.orderRepo.Create(order)
}

func (u *orderUsecase) GetOrderByID(id int64) (*domain.Order, error) {
	if id <= 0 {
		return nil, errors.New("invalid order id")
	}
	
	return u.orderRepo.FindByID(id)
}

func (u *orderUsecase) GetOrdersByUserID(userID int64) ([]*domain.Order, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}
	
	return u.orderRepo.FindByUserID(userID)
}

func (u *orderUsecase) UpdateOrder(order *domain.Order) error {
	if order.ID <= 0 {
		return errors.New("invalid order id")
	}
	
	existingOrder, err := u.orderRepo.FindByID(order.ID)
	if err != nil {
		return err
	}
	if existingOrder == nil {
		return errors.New("order not found")
	}
	
	return u.orderRepo.Update(order)
}

func (u *orderUsecase) DeleteOrder(id int64) error {
	if id <= 0 {
		return errors.New("invalid order id")
	}
	
	existingOrder, err := u.orderRepo.FindByID(id)
	if err != nil {
		return err
	}
	if existingOrder == nil {
		return errors.New("order not found")
	}
	
	return u.orderRepo.Delete(id)
}