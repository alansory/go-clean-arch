package controller

import (
	"go-esb-test/internal/model"
	"go-esb-test/internal/usecase"
	"go-esb-test/internal/util"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	UseCase *usecase.UserUseCase
	Log     *logrus.Logger
}

func NewUserController(useCase *usecase.UserUseCase, log *logrus.Logger) *UserController {
	return &UserController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *UserController) List(ctx *fiber.Ctx) error {

	request := &model.SearchUserRequest{
		Fullname: ctx.Query("fullname", ""),
		Username: ctx.Query("username", ""),
		Email:    ctx.Query("email", ""),
		Page:     ctx.QueryInt("page", 1),
		PerPage:  ctx.QueryInt("per_page", 10),
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

	return util.SuccessResponse("User list retrieved successfully.", responses, paging, ctx)
}
