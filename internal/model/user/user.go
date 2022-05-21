package user

import "time"

// User Entity -.
type User struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Age       int64     `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
