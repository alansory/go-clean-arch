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

	// Determine the status code based on the error type
	switch e := err.(type) {
	case *fiber.Error:
		statusCode = e.Code
	case validator.ValidationErrors:
		statusCode = fiber.StatusUnprocessableEntity
	default:
		statusCode = fiber.StatusInternalServerError
	}

	report := fiber.Map{
		"error": fiber.Map{
			"status_code": statusCode,
			"message":     err.Error(),
		},
	}

	switch v := err.(type) {
	case validator.ValidationErrors:
		statusCode = fiber.StatusUnprocessableEntity
		report["error"].(fiber.Map)["status_code"] = statusCode
		report["error"].(fiber.Map)["message"] = "422 Unprocessable Entity"
		errors := make(map[string]string)

		for _, validationErr := range v {
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
		report["error"].(fiber.Map)["errors"] = errors
	case error:
		if err.Error() == "EOF" {
			statusCode = fiber.StatusBadRequest
			report["error"].(fiber.Map)["status_code"] = statusCode
			report["error"].(fiber.Map)["message"] = "JSON data is missing or malformed"
		}
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
