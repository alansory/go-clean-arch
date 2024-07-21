package converter

import (
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"
)

func InvoiceToResponse(invoice *entity.Invoice) *model.InvoiceResponse {
	var subTotal float64
	for _, item := range invoice.InvoiceItems {
		subTotal += float64(item.Quantity) * item.UnitPrice
	}

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
		TotalItems:      len(invoice.InvoiceItems),
		SubTotal:        subTotal,
	}
}
