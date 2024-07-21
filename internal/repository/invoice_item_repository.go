package repository

import (
	"go-esb-test/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (r *Repository[T]) FindByIdAndInvoiceId(db *gorm.DB, entity *T, id int64, invoiceId int64) error {
	query := db.Where("id = ?", id).Where("invoice_id = ?", invoiceId)
	return query.Take(entity).Error
}
