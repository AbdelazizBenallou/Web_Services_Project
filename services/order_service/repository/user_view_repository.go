package repository

type UserViewRepository interface {
	Exists(userID int64) (bool, error)
}

