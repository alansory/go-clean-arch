package converter

import (
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"
)

func InvoiceToResponse(invoice *entity.Invoice) *model.InvoiceResponse {
	return &model.InvoiceResponse{
		ID:              invoice.ID,
		InvoiceNumber:   invoice.InvoiceNumber,
		InvoiceSubject:  invoice.InvoiceSubject,
		IssueDate:       invoice.IssueDate,
		DueDate:         invoice.DueDate,
		Status:          invoice.Status,
		CustomerID:      invoice.CustomerID,
		CustomerName:    invoice.CustomerName,
		CustomerAddress: invoice.CustomerAddress,
		CreatedAt:       invoice.CreatedAt,
		UpdatedAt:       invoice.UpdatedAt,
		Customer:        UserToResponse(&invoice.Customer),
		InvoiceItems:    InvoiceItemsToResponse(invoice.InvoiceItems),
	}
}
