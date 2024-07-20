package converter

import (
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"
)

func ItemToResponse(item *entity.Item) *model.ItemResponse {
	return &model.ItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Quantity:    item.Quantity,
		UnitPrice:   item.UnitPrice,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func ItemsToResponse(items []entity.Item) []*model.ItemResponse {
	var responses []*model.ItemResponse
	for _, item := range items {
		responses = append(responses, ItemToResponse(&item))
	}
	return responses
}
