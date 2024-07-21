package usecase

import (
	"context"
	"go-esb-test/internal/model"
	"go-esb-test/internal/model/converter"
	"go-esb-test/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ItemUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	ItemRepository *repository.ItemRepository
}

func NewItemUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	ItemRepository *repository.ItemRepository,
) *ItemUseCase {
	return &ItemUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		ItemRepository: ItemRepository,
	}
}

func (c *ItemUseCase) List(ctx context.Context, request *model.SearchItemRequest) ([]model.ItemResponse, int64, error) {
	items, total, err := c.ItemRepository.Search(c.DB, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting items")
		return nil, 0, err
	}

	responses := make([]model.ItemResponse, len(items))
	for i, item := range items {
		responses[i] = *converter.ItemToResponse(&item)
	}

	return responses, total, nil
}
