package repository

import (
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	Repository[entity.Invoice]
	Log *logrus.Logger
}

func NewInvoiceRepository(log *logrus.Logger) *InvoiceRepository {
	return &InvoiceRepository{
		Log: log,
	}
}

func (r *InvoiceRepository) Search(db *gorm.DB, request *model.SearchInvoiceRequest) ([]entity.Invoice, int64, error) {
	var invoices []entity.Invoice

	query := db.Scopes(r.FilterInvoice(request)).
		Offset((request.Page - 1) * request.PerPage).
		Limit(request.PerPage).
		Preload("InvoiceItems.Item").
		Preload("Customer")

	if request.TotalItems > 0 {
		subQuery := `
			SELECT invoice_id
			FROM invoice_items
			GROUP BY invoice_id
			HAVING COUNT(*) = ?
		`
		// Join the raw SQL subquery with the invoices
		query = query.Joins(
			"JOIN (?) AS subquery ON invoices.id = subquery.invoice_id",
			db.Raw(subQuery, request.TotalItems).Table("(?)"),
		)
	}

	if err := query.Find(&invoices).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	totalQuery := db.Model(&entity.Invoice{}).
		Scopes(r.FilterInvoice(request))

	if request.TotalItems > 0 {
		totalSubQuery := `
				SELECT invoice_id
				FROM invoice_items
				GROUP BY invoice_id
				HAVING COUNT(*) = ?
			`
		totalQuery = totalQuery.Joins(
			"JOIN (?) AS subquery ON invoices.id = subquery.invoice_id",
			db.Raw(totalSubQuery, request.TotalItems).Table("(?)"),
		)
	}

	if err := totalQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return invoices, total, nil
}

func (r *InvoiceRepository) FilterInvoice(request *model.SearchInvoiceRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {

		if invoiceNumber := request.InvoiceNumber; invoiceNumber != "" {
			invoiceNumber = "%" + invoiceNumber + "%"
			tx = tx.Where("invoice_number LIKE ?", invoiceNumber)
		}

		if issueDate := request.IssueDate; issueDate != "" {
			tx = tx.Where("issue_date = ?", request.IssueDate)
		}

		if dueDate := request.DueDate; dueDate != "" {
			tx = tx.Where("due_date = ?", request.IssueDate)
		}

		if subject := request.InvoiceSubject; subject != "" {
			subject = "%" + subject + "%"
			tx = tx.Where("subject LIKE ?", subject)
		}

		if name := request.CustomerName; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("cutomer_name LIKE ?", name)
		}

		if status := request.Status; status != "" {
			tx = tx.Where("status = ?", status)
		}

		return tx
	}
}
