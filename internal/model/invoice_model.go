package model

import "time"

type InvoiceResponse struct {
	ID              int64                  `json:"id"`
	InvoiceNumber   string                 `json:"invoice_number"`
	InvoiceSubject  string                 `json:"invoice_subject"`
	IssueDate       *time.Time             `json:"issue_date"`
	DueDate         *time.Time             `json:"due_date"`
	Status          string                 `json:"status"`
	CustomerID      int64                  `json:"customer_id"`
	CustomerName    string                 `json:"customer_name"`
	CustomerAddress string                 `json:"customer_address"`
	CreatedAt       *time.Time             `json:"created_at"`
	UpdatedAt       *time.Time             `json:"updated_at"`
	Customer        *UserResponse          `json:"customer"`
	InvoiceItems    []*InvoiceItemResponse `json:"invoice_items"`
	TotalItems      int                    `json:"total_items"`
	SubTotal        float64                `json:"sub_total"`
}

type InvoiceRequest struct {
	InvoiceSubject string        `json:"invoice_subject" validate:"required,max=150"`
	IssueDate      *time.Time    `json:"issue_date" validate:"required"`
	DueDate        *time.Time    `json:"due_date" validate:"required"`
	CustomerID     int64         `json:"customer_id" validate:"required"`
	Status         string        `json:"status" validate:"required,oneof=paid unpaid"`
	Items          []ItemRequest `json:"items" validate:"required,dive,required"`
}

type SearchInvoiceRequest struct {
	InvoiceNumber  string `json:"invoice_number"`
	InvoiceSubject string `json:"invoice_subject"`
	IssueDate      string `json:"issue_date"`
	DueDate        string `json:"due_date"`
	Status         string `json:"status"`
	CustomerName   string `json:"customer_name"`
	TotalItems     int    `json:"total_items"`
	Page           int    `json:"page"`
	PerPage        int    `json:"per_page"`
}
