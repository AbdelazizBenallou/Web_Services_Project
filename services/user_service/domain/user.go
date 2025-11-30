// Package domain contains the core business entities.
package domain

import (
	"time"
)

// User represents a user entity in the system.
type User struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
