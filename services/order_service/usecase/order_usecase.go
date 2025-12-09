package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"order_service/domain"
)

type orderUsecase struct {
	orderRepo domain.OrderRepository
}

func NewOrderUsecase(orderRepo domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{orderRepo: orderRepo}
}

// CreateOrder ينشئ طلب جديد مع التحقق من المستخدم والمنتج
func (u *orderUsecase) CreateOrder(order *domain.Order) error {
	// التحقق من المدخلات الأساسية
	if order.UserID <= 0 {
		return errors.New("user_id is required")
	}
	if order.ProductID <= 0 {
		return errors.New("product_id is required")
	}
	if order.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// التحقق من المستخدم في user_service
	user, err := u.getUserByID(order.UserID)
	if err != nil {
		return fmt.Errorf("user validation failed: %v", err)
	}

	// التحقق من المنتج في product_service
	product, err := u.getProductByID(order.ProductID)
	if err != nil {
		return fmt.Errorf("product validation failed: %v", err)
	}

	// حساب السعر الكلي للطلب
	order.TotalPrice = float64(order.Quantity) * product.Price

	// تعيين الحالة الافتراضية إذا لم تحدد
	if order.Status == "" {
		order.Status = "pending"
	}

	// إنشاء الطلب في قاعدة البيانات
	return u.orderRepo.Create(order)
}

// GetOrderByID يسترجع طلب محدد
func (u *orderUsecase) GetOrderByID(id int64) (*domain.Order, error) {
	if id <= 0 {
		return nil, errors.New("invalid order id")
	}

	return u.orderRepo.FindByID(id)
}

// GetOrdersByUserID يسترجع جميع الطلبات لمستخدم محدد
func (u *orderUsecase) GetOrdersByUserID(userID int64) ([]*domain.Order, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	return u.orderRepo.FindByUserID(userID)
}

// UpdateOrder لتحديث بيانات طلب موجود
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

// DeleteOrder لحذف طلب موجود
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

//////////////////////
// دوال داخلية للتواصل مع الخدمات الأخرى
//////////////////////

// getUserByID يتواصل مع user_service للتحقق من المستخدم
func (u *orderUsecase) getUserByID(userID int64) (*domain.User, error) {
	url := fmt.Sprintf("http://user_service:8080/api/v1/users/%d", userID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("user not found")
	}

	var user domain.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// getProductByID يتواصل مع product_service للتحقق من المنتج والحصول على السعر
func (u *orderUsecase) getProductByID(productID int64) (*domain.Product, error) {
	url := fmt.Sprintf("http://product_service:8080/api/v1/products/%d", productID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("product not found")
	}

	var product domain.Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}
