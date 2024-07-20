package model

import "time"

type ItemResponse struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Quantity    int        `json:"quantity"`
	UnitPrice   float32    `json:"unit_price"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type ItemRequest struct {
	ID        int64   `json:"id" validate:"required"`
	Name      string  `json:"name" validate:"required,max=255"`
	Quantity  int     `json:"quantity" validate:"required"`
	UnitPrice float32 `json:"unit_price" validate:"required"`
}
