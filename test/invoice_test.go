package test

import (
	"encoding/json"
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	tx := db.Begin()
	assert.Nil(t, tx.Error)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test panicked: %v", r)
		}
		err := tx.Rollback().Error
		assert.Nil(t, err, "Rollback failed")
	}()

	user := new(entity.User)
	item := new(entity.Item)
	CreateUsers(user, 1)
	CreateItems(item, 1)
	createdUser := GetFirstUser(t)
	createdItem := GetFirstItem(t)

	issueDate, err := time.Parse(time.RFC3339, "2024-08-01T00:00:00Z")
	assert.Nil(t, err)
	dueDate, err := time.Parse(time.RFC3339, "2024-08-15T00:00:00Z")
	assert.Nil(t, err)

	requestBody := model.InvoiceRequest{
		InvoiceSubject: "Service data 2",
		IssueDate:      &issueDate,
		DueDate:        &dueDate,
		CustomerID:     createdUser.ID,
		Status:         "unpaid",
		Items: []model.ItemRequest{
			{
				ItemID:    createdItem.ID,
				Name:      "Service A",
				Quantity:  2,
				UnitPrice: 100.00,
			},
		},
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/invoices", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.SuccessResponse)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	invoiceData, ok := responseBody.Data.(map[string]interface{})
	assert.True(t, ok, "Failed to assert responseBody.Data as map[string]interface{}")

	assert.Equal(t, requestBody.InvoiceSubject, invoiceData["invoice_subject"])
	assert.Equal(t, requestBody.Status, invoiceData["status"])
	assert.NotNil(t, invoiceData["created_at"])
	assert.NotNil(t, invoiceData["updated_at"])
	assert.NotNil(t, invoiceData["id"])
}

func TestCreateInvoiceFailed(t *testing.T) {
	tx := db.Begin()
	assert.Nil(t, tx.Error)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test panicked: %v", r)
		}
		err := tx.Rollback().Error
		assert.Nil(t, err, "Rollback failed")
	}()

	user := new(entity.User)
	item := new(entity.Item)
	CreateUsers(user, 1)
	CreateItems(item, 1)
	createdUser := GetFirstUser(t)
	createdItem := GetFirstItem(t)

	issueDate, err := time.Parse(time.RFC3339, "2024-08-01T00:00:00Z")
	assert.Nil(t, err)
	dueDate, err := time.Parse(time.RFC3339, "2024-08-15T00:00:00Z")
	assert.Nil(t, err)

	// Example of invalid request: Missing InvoiceSubject
	requestBody := model.InvoiceRequest{
		IssueDate:  &issueDate,
		DueDate:    &dueDate,
		CustomerID: createdUser.ID,
		Status:     "unpaid",
		Items: []model.ItemRequest{
			{
				ItemID:    createdItem.ID,
				Name:      "Service A",
				Quantity:  2,
				UnitPrice: 100.00,
			},
		},
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/invoices", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.ErrorResponse)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnprocessableEntity, response.StatusCode)

	assert.Contains(t, responseBody.Error.Message, "422 Unprocessable Entity")
	assert.NotEmpty(t, responseBody.Error.Errors)
	assert.Contains(t, responseBody.Error.Errors, "invoice_subject")
}

func TestGetInvoice(t *testing.T) {
	tx := db.Begin()
	assert.Nil(t, tx.Error)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test panicked: %v", r)
		}
		err := tx.Rollback().Error
		assert.Nil(t, err, "Rollback failed")
	}()

	CreateInvoices(tx, 1, t)
	createdInvoice := GetFirstInvoice(t)

	invoiceID := strconv.FormatInt(createdInvoice.ID, 10)
	request := httptest.NewRequest(http.MethodGet, "/invoices/"+invoiceID, nil)
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.SuccessResponse)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	invoiceData, ok := responseBody.Data.(map[string]interface{})
	assert.True(t, ok, "Failed to assert responseBody.Data as map[string]interface{}")

	assert.NotNil(t, invoiceData["created_at"])
	assert.NotNil(t, invoiceData["updated_at"])
	assert.NotNil(t, invoiceData["id"])
}

func TestGetInvoiceFailed(t *testing.T) {
	tx := db.Begin()
	assert.Nil(t, tx.Error)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test panicked: %v", r)
		}
		err := tx.Rollback().Error
		assert.Nil(t, err, "Rollback failed")
	}()

	CreateInvoices(tx, 1, t)
	createdInvoice := GetFirstInvoice(t)

	invoiceID := strconv.FormatInt(createdInvoice.ID, 10)
	invalidURL := "/invoices/" + invoiceID + "/invalidSegment"
	request := httptest.NewRequest(http.MethodGet, invalidURL, nil)
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}
