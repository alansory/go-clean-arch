package controller

import (
	"go-esb-test/internal/model"
	"go-esb-test/internal/usecase"
	"go-esb-test/internal/util"
	"math"
	"strconv"

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

func (c *InvoiceController) List(ctx *fiber.Ctx) error {

	request := &model.SearchInvoiceRequest{
		InvoiceSubject: ctx.Query("subject", ""),
		InvoiceNumber:  ctx.Query("invoice_number", ""),
		IssueDate:      ctx.Query("issue_date", ""),
		DueDate:        ctx.Query("due_date", ""),
		CustomerName:   ctx.Query("name", ""),
		Status:         ctx.Query("status", ""),
		TotalItems:     ctx.QueryInt("total_items", 0),
		Page:           ctx.QueryInt("page", 1),
		PerPage:        ctx.QueryInt("per_page", 10),
	}

	responses, total, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching invoice")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		PerPage:   request.PerPage,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.PerPage))),
	}

	return util.SuccessResponse("Invoice list retrieved successfully.", responses, paging, ctx)
}

func (c *InvoiceController) Get(ctx *fiber.Ctx) error {
	response, err := c.UseCase.Get(ctx)
	if err != nil {
		c.Log.WithError(err).Error("error getting invoice")
		return util.ErrorResponse(err, ctx)
	}

	return util.SuccessResponse("Invoice details retrieved successfully.", response, nil, ctx)
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

	return util.SuccessResponse("Invoice created successfully..", response, nil, ctx)
}

func (c *InvoiceController) Update(ctx *fiber.Ctx) error {

	request := new(model.InvoiceRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return util.ErrorResponse(err, ctx)
	}

	idStr := ctx.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.Log.WithError(err).Error("invalid id")
		return util.ErrorResponse(err, ctx)
	}

	response, err := c.UseCase.Update(id, ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating invoice")
		return util.ErrorResponse(err, ctx)
	}

	return util.SuccessResponse("Invoice updated successfully..", response, nil, ctx)
}

func (c *InvoiceController) Delete(ctx *fiber.Ctx) error {
	_, err := c.UseCase.Delete(ctx)
	if err != nil {
		c.Log.WithError(err).Error("error deleting invoice")
		return util.ErrorResponse(err, ctx)
	}

	return util.SuccessResponse("Invoice deleted successfully.", nil, nil, ctx)
}
