package model

import "time"

type InvoiceItemResponse struct {
	ID        int64         `json:"id"`
	InvoiceID int64         `json:"invoice_id"`
	ItemID    int64         `json:"item_id"`
	ItemName  string        `json:"item_name"`
	Quantity  int           `json:"quantity"`
	UnitPrice float64       `json:"unit_price"`
	CreatedAt *time.Time    `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at"`
	Item      *ItemResponse `json:"item"`
}
