package model

import "time"

// User represents a user entity
type User struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Password  string
}
