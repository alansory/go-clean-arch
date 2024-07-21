package entity

import "time"

type InvoiceItem struct {
	ID        int64      `gorm:"column:id;primaryKey"`
	InvoiceID int64      `gorm:"column:invoice_id"`
	ItemID    int64      `gorm:"column:item_id"`
	ItemName  string     `gorm:"column:item_name"`
	Quantity  int        `gorm:"column:quantity"`
	UnitPrice float64    `gorm:"column:unit_price"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	Invoice   Invoice    `gorm:"foreignKey:invoice_id;references:id"`
	Item      Item       `gorm:"foreignKey:item_id;references:id"`
}

func (a *InvoiceItem) TableName() string {
	return "invoice_items"
}
