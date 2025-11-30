// repository/user_repository.go
package repository

import "user_services/domain"

type UserRepository interface {
	GetByID(id int64) (*domain.User, error)
	GetAll() ([]*domain.User,  error)
}
