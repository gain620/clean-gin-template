package user

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Age       int64     `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
