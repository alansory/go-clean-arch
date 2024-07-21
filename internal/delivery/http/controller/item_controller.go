package controller

import (
	"go-esb-test/internal/model"
	"go-esb-test/internal/usecase"
	"go-esb-test/internal/util"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ItemController struct {
	UseCase *usecase.ItemUseCase
	Log     *logrus.Logger
}

func NewItemController(useCase *usecase.ItemUseCase, log *logrus.Logger) *ItemController {
	return &ItemController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *ItemController) List(ctx *fiber.Ctx) error {

	request := &model.SearchItemRequest{
		Name:        ctx.Query("name", ""),
		Description: ctx.Query("description", ""),
		Page:        ctx.QueryInt("page", 1),
		PerPage:     ctx.QueryInt("per_page", 10),
	}

	responses, total, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching user")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		PerPage:   request.PerPage,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.PerPage))),
	}

	return util.SuccessResponse("Item list retrieved successfully.", responses, paging, ctx)
}
