package controller

import (
	"go-esb-test/internal/model"
	"go-esb-test/internal/usecase"
	"go-esb-test/internal/util"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type InvoiceController struct {
	UseCase *usecase.InvoiceUseCase
	Log     *logrus.Logger
}

func NewInvoiceController(useCase *usecase.InvoiceUseCase, log *logrus.Logger) *InvoiceController {
	return &InvoiceController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *InvoiceController) Create(ctx *fiber.Ctx) error {
	request := new(model.InvoiceRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return util.ErrorResponse(err, ctx)
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create invoice")
		return util.ErrorResponse(err, ctx)
	}

	return util.SuccessResponse("Created.", response, ctx)
}
