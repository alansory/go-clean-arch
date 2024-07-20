package usecase

import (
	"context"
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"
	"go-esb-test/internal/model/converter"
	"go-esb-test/internal/repository"
	"go-esb-test/internal/util"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InvoiceUseCase struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	Validate              *validator.Validate
	InvoiceRepository     *repository.InvoiceRepository
	UserRepository        *repository.UserRepository
	InvoiceItemRepository *repository.InvoiceItemRepository
}

func NewInvoiceUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	invoiceRepository *repository.InvoiceRepository,
	userRepository *repository.UserRepository,
	invoiceItemRepository *repository.InvoiceItemRepository,
) *InvoiceUseCase {
	return &InvoiceUseCase{
		DB:                    db,
		Log:                   logger,
		Validate:              validate,
		InvoiceRepository:     invoiceRepository,
		UserRepository:        userRepository,
		InvoiceItemRepository: invoiceItemRepository,
	}
}

func (c *InvoiceUseCase) Create(ctx context.Context, request *model.InvoiceRequest) (*model.InvoiceResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if request.Status == "" {
		request.Status = "unpaid"
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, err
	}

	invoiceNumber, err := util.GenerateInvoiceNumber("ESB", 3)
	if err != nil {
		c.Log.WithError(err).Error("error generating invoice number")
		return nil, err
	}

	customer := new(entity.User)
	if err := c.UserRepository.FindById(tx, customer, request.CustomerID); err != nil {
		c.Log.Warnf("Failed find customer by customer_id : %+v", err)
		return nil, err
	}

	invoice := &entity.Invoice{
		InvoiceNumber:   invoiceNumber,
		InvoiceSubject:  request.InvoiceSubject,
		CustomerID:      customer.ID,
		CustomerName:    customer.Fullname,
		CustomerAddress: customer.Address,
		IssueDate:       request.IssueDate,
		DueDate:         request.DueDate,
		Status:          request.Status,
	}

	if err := c.InvoiceRepository.Create(tx, invoice); err != nil {
		c.Log.WithError(err).Error("failed to create invoice")
		return nil, err
	}

	for _, item := range request.Items {
		invoiceItem := &entity.InvoiceItem{
			InvoiceID: invoice.ID,
			ItemID:    item.ID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		}

		if err := c.InvoiceItemRepository.Create(tx, invoiceItem); err != nil {
			c.Log.WithError(err).Error("failed to create invoice item")
			return nil, err
		}
	}

	if err := c.InvoiceRepository.FindById(tx, invoice, invoice.ID, "Customer", "InvoiceItems.Item"); err != nil {
		c.Log.WithError(err).Error("failed to reload invoice with items")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return converter.InvoiceToResponse(invoice), nil
}

func (c *InvoiceUseCase) Get(ctx *fiber.Ctx) (*model.InvoiceResponse, error) {
	idStr := ctx.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.Log.WithError(err).Error("invalid id")
		return nil, err
	}

	invoice := new(entity.Invoice)
	if err := c.InvoiceRepository.FindById(c.DB, invoice, id, "Customer", "InvoiceItems.Item"); err != nil {
		c.Log.WithError(err).Error("error getting invoice")
		return nil, err
	}

	return converter.InvoiceToResponse(invoice), nil
}

func (c *InvoiceUseCase) Delete(ctx *fiber.Ctx) (*model.InvoiceResponse, error) {
	idStr := ctx.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.Log.WithError(err).Error("invalid id")
		return nil, err
	}

	invoice := new(entity.Invoice)
	if err := c.InvoiceRepository.Delete(c.DB, invoice, id); err != nil {
		c.Log.WithError(err).Error("error deleting invoice")
		return nil, err
	}

	return converter.InvoiceToResponse(invoice), nil
}
