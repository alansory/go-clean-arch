package entity

import "time"

type Item struct {
	ID          int64      `gorm:"column:id;primaryKey"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	Quantity    int        `gorm:"column:quantity"`
	UnitPrice   float32    `gorm:"column:unit_price"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (a *Item) TableName() string {
	return "invoice_items"
}
