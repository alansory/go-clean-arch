package util

import (
	"fmt"
	"go-esb-test/internal/model"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(message string, data interface{}, paging *model.PageMetadata, ctx *fiber.Ctx) error {
	statusCode := fiber.StatusOK

	result := model.SuccessResponse{
		Data:       data,
		Paging:     paging,
		Message:    message,
		StatusCode: statusCode,
	}

	return ctx.Status(statusCode).JSON(result)
}

func ErrorResponse(err error, ctx *fiber.Ctx) error {
	var statusCode int
	var errorMessage string
	var errors map[string]string

	// Determine the status code based on the error type
	switch e := err.(type) {
	case *fiber.Error:
		statusCode = e.Code
		errorMessage = e.Message
	case validator.ValidationErrors:
		statusCode = fiber.StatusUnprocessableEntity
		errorMessage = "422 Unprocessable Entity"
		errors = make(map[string]string)
		for _, validationErr := range e {
			field := toSnakeCase(validationErr.Field())
			switch validationErr.Tag() {
			case "required":
				errors[strings.ToLower(field)] = fmt.Sprintf("The %s field is required", field)
			case "email":
				errors[field] = fmt.Sprintf("The %s field is not a valid email", field)
			case "oneof":
				errors[field] = fmt.Sprintf("The %s field must be one of: %s", field, validationErr.Param())
			case "gte":
				errors[field] = fmt.Sprintf("The %s field value must be greater than %s", field, validationErr.Param())
			case "lte":
				errors[field] = fmt.Sprintf("The %s field value must be lower than %s", field, validationErr.Param())
			}
		}
	default:
		statusCode = fiber.StatusInternalServerError
		errorMessage = "Internal Server Error"
	}

	if err.Error() == "EOF" {
		statusCode = fiber.StatusBadRequest
		errorMessage = "JSON data is missing or malformed"
	}

	report := model.ErrorResponse{
		Error: struct {
			StatusCode int               `json:"status_code"`
			Message    string            `json:"message"`
			Errors     map[string]string `json:"errors,omitempty"`
		}{
			StatusCode: statusCode,
			Message:    errorMessage,
			Errors:     errors,
		},
	}

	return ctx.Status(statusCode).JSON(report)
}

func toSnakeCase(str string) string {
	var sb strings.Builder
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if i > 0 && runes[i] >= 'A' && runes[i] <= 'Z' {
			// Check for consecutive uppercase letters
			if i+1 < len(runes) && runes[i+1] >= 'A' && runes[i+1] <= 'Z' {
				sb.WriteRune('_')
			} else if i > 0 && runes[i-1] >= 'a' && runes[i-1] <= 'z' {
				sb.WriteRune('_')
			}
		}
		sb.WriteRune(runes[i])
	}
	return strings.ToLower(sb.String())
}
