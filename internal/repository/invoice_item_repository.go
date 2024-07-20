package repository

import (
	"go-esb-test/internal/entity"

	"github.com/sirupsen/logrus"
)

type InvoiceItemRepository struct {
	Repository[entity.InvoiceItem]
	Log *logrus.Logger
}

func NewInvoiceItemRepository(log *logrus.Logger) *InvoiceItemRepository {
	return &InvoiceItemRepository{
		Log: log,
	}
}
