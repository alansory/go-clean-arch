package test

import (
	"errors"
	"fmt"
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var idNotNull = "id is not null"

func ClearAll() {
	ClearUsers()
	ClearItems()
	ClearInvoices()
	ClearInvoiceItems()
}

func ClearUsers() {
	err := db.Where(idNotNull).Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func ClearItems() {
	err := db.Where(idNotNull).Delete(&entity.Item{}).Error
	if err != nil {
		log.Fatalf("Failed clear item data : %+v", err)
	}
}

func ClearInvoiceItems() {
	err := db.Where(idNotNull).Delete(&entity.InvoiceItem{}).Error
	if err != nil {
		log.Fatalf("Failed clear invoice item data : %+v", err)
	}
}

func ClearInvoices() {
	err := db.Where(idNotNull).Delete(&entity.Invoice{}).Error
	if err != nil {
		log.Fatalf("Failed clear invoice data : %+v", err)
	}
}

func CreateUsers(user *entity.User, total int) {
	for i := 0; i < total; i++ {
		// Generate unique email address
		email := fmt.Sprintf("user%d@example.com", i)

		// Check if the email already exists
		err := db.Where("email = ?", email).First(&user).Error
		if err == nil {
			// If the email exists, skip creation
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle other errors
			log.Fatalf("Failed to check if user exists: %+v", err)
		}

		// Create a new user
		user := &entity.User{
			Fullname: "User",
			Username: "user" + strconv.Itoa(i),
			Email:    email,
			Phone:    "08000000" + strconv.Itoa(i),
			Address:  "Jalan Belum Jadi",
		}
		err = db.Create(user).Error
		if err != nil {
			log.Fatalf("Failed to create user data: %+v", err)
		}
	}
}

func CreateItems(item *entity.Item, total int) {
	for i := 0; i < total; i++ {
		item := &entity.Item{
			Name:        "Item",
			Type:        strconv.Itoa(i),
			Description: "Description" + strconv.Itoa(i),
		}
		err := db.Create(item).Error
		if err != nil {
			log.Fatalf("Failed create item data : %+v", err)
		}
	}
}

func CreateInvoices(tx *gorm.DB, total int, t *testing.T) {
	for i := 0; i < total; i++ {
		issueDate := time.Now().AddDate(0, 0, -i)
		dueDate := issueDate.AddDate(0, 1, 0)

		user := new(entity.User)
		item := new(entity.Item)
		CreateUsers(user, 1)
		CreateItems(item, 1)
		createdUser := GetFirstUser(t)
		createdItem := GetFirstItem(t)

		// Create InvoiceRequest object
		requestBody := model.InvoiceRequest{
			IssueDate:  &issueDate,
			DueDate:    &dueDate,
			CustomerID: createdUser.ID,
			Status:     "unpaid",
			Items: []model.ItemRequest{
				{
					ItemID:    createdItem.ID,
					Name:      fmt.Sprintf("Service A %d", i),
					Quantity:  2,
					UnitPrice: 100.00,
				},
			},
		}

		// Map the requestBody to an Invoice entity
		invoice := &entity.Invoice{
			InvoiceSubject: requestBody.InvoiceSubject,
			IssueDate:      requestBody.IssueDate,
			DueDate:        requestBody.DueDate,
			CustomerID:     requestBody.CustomerID,
			Status:         requestBody.Status,
		}

		// Create the invoice record
		err := tx.Create(invoice).Error
		if err != nil {
			log.Fatalf("Failed to create invoice: %+v", err)
		}

		for _, itemReq := range requestBody.Items {
			item := &entity.InvoiceItem{
				InvoiceID: invoice.ID,
				ItemID:    itemReq.ItemID,
				ItemName:  itemReq.Name,
				Quantity:  itemReq.Quantity,
				UnitPrice: itemReq.UnitPrice,
			}
			err = tx.Create(item).Error
			if err != nil {
				log.Fatalf("Failed to create item: %+v", err)
			}
		}
	}
}

func GetFirstUser(t *testing.T) *entity.User {
	user := new(entity.User)
	err := db.Order("id desc").First(user).Error
	assert.Nil(t, err)
	return user
}

func GetFirstItem(t *testing.T) *entity.Item {
	item := new(entity.Item)
	err := db.Order("id desc").First(item).Error
	assert.Nil(t, err)
	return item
}

func GetFirstInvoice(t *testing.T) *entity.Invoice {
	invoice := new(entity.Invoice)
	err := db.Order("id desc").First(invoice).Error
	assert.Nil(t, err)
	return invoice
}
