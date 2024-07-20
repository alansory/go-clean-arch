package model

import "time"

type UserResponse struct {
	ID        int64      `json:"id"`
	Fullname  string     `json:"fullname"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
