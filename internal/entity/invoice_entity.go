package entity

import "time"

type Invoice struct {
	ID              int64         `gorm:"column:id;primaryKey"`
	InvoiceNumber   string        `gorm:"column:invoice_number"`
	InvoiceSubject  string        `gorm:"column:invoice_subject"`
	IssueDate       *time.Time    `gorm:"column:issue_date"`
	DueDate         *time.Time    `gorm:"column:due_date"`
	Status          string        `gorm:"column:status"`
	CustomerID      int64         `gorm:"column:customer_id"`
	CustomerName    string        `gorm:"column:customer_name"`
	CustomerAddress string        `gorm:"column:customer_address"`
	CreatedAt       *time.Time    `gorm:"column:created_at"`
	UpdatedAt       *time.Time    `gorm:"column:updated_at"`
	Customer        User          `gorm:"foreignKey:customer_id;references:id"`
	InvoiceItems    []InvoiceItem `gorm:"foreignKey:invoice_id;references:id"`
}

func (a *Invoice) TableName() string {
	return "invoices"
}
