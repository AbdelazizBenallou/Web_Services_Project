package domain

import (
	"time"
)

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleWorker Role = "worker"
	RoleClient Role = "client"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleWorker, RoleClient:
		return true
	default:
		return false
	}
}

type Profile struct {
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	Address   string    `json:"address,omitempty"`
}

type User struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`

	Role    Role    `json:"role"`
	Profile Profile `json:"profile,omitempty"`
}
