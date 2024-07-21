package model

import "time"

type ItemResponse struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type ItemRequest struct {
	ID        int64   `json:"id"`
	ItemID    int64   `json:"item_id" validate:"required"`
	Name      string  `json:"name" validate:"required,max=255"`
	Quantity  int     `json:"quantity" validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required"`
}

type SearchItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Page        int    `json:"page"`
	PerPage     int    `json:"per_page"`
}
