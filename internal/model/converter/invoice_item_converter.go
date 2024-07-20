package converter

import (
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"
)

func InvoiceItemToResponse(item *entity.InvoiceItem) *model.InvoiceItemResponse {
	return &model.InvoiceItemResponse{
		ID:        item.ID,
		InvoiceID: item.InvoiceID,
		ItemID:    item.ItemID,
		Quantity:  item.Quantity,
		UnitPrice: item.UnitPrice,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
		Item:      ItemToResponse(&item.Item),
	}
}

func InvoiceItemsToResponse(items []entity.InvoiceItem) []*model.InvoiceItemResponse {
	var responses []*model.InvoiceItemResponse
	for _, item := range items {
		responses = append(responses, InvoiceItemToResponse(&item))
	}
	return responses
}
