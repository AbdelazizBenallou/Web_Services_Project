package usecase

import (
	"errors"
	"order_service/domain"
	"order_service/repository"
	"time"
	"encoding/json"
)

type OrderUseCase interface {
	CreateOrder(userID int64, items []domain.OrderItem) (*domain.Order, error)
}

type orderUseCase struct {
	orderRepo repository.OrderRepository
	userViewRepo repository.UserViewRepository
	publisher    EventPublisher
}

func NewOrderUseCase(orderRepo repository.OrderRepository,
		     userViewRepo repository.UserViewRepository,
	             publisher EventPublisher,
		) OrderUseCase {
			return &orderUseCase{
			orderRepo:    orderRepo,
			userViewRepo: userViewRepo,
			publisher:    publisher,
	}
}

func (uc *orderUseCase) CreateOrder(
	userID int64,
	items []domain.OrderItem,
) (*domain.Order, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	exists, err := uc.userViewRepo.Exists(userID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("user not registered in order service")
	}

	if len(items) == 0 {
		return nil, errors.New("order must have items")
	}

	order := &domain.Order{
		UserID:    userID,
		Status:    domain.StatusPendingInventory,
		Items:     items,
		CreatedAt: time.Now(),
	}

	if err := uc.orderRepo.Create(order); err != nil {
		return nil, err
	}

	if uc.publisher != nil {
		event := map[string]interface{}{
			"order_id":   order.ID,
			"user_id":    order.UserID,
			"status":     order.Status,
			"items":      order.Items,
			"created_at": order.CreatedAt,
		}

		data, _ := json.Marshal(event)
		_ = uc.publisher.Publish("order.created", data)
	}

	return order, nil
}


