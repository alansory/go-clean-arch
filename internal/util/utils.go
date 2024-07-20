package util

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

func GenerateRandomString(length int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	_, err := rand.Read(result)
	if err != nil {
		return "", err
	}

	for i, b := range result {
		result[i] = chars[b%byte(len(chars))]
	}

	return string(result), nil
}

func GenerateInvoiceNumber(prefix string, randomLength int) (string, error) {
	// Format the current data as "ddmmyy"
	now := time.Now()
	dateStr := now.Format("020106")

	// Generate a random string length 3 or more
	randomString, err := GenerateRandomString(randomLength)
	if err != nil {
		return "", err
	}

	// Combine everything into the final invoice number
	invoiceNumber := fmt.Sprintf("%s-%s-%s", prefix, dateStr, randomString)
	return strings.ToUpper(invoiceNumber), nil
}
