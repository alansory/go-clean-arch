package entity

import "time"

type Item struct {
	ID          int64      `gorm:"column:id;primaryKey"`
	Name        string     `gorm:"column:name"`
	Type        string     `gorm:"column:type"`
	Description string     `gorm:"column:description"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (a *Item) TableName() string {
	return "items"
}
