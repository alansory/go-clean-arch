package main

import (
	"go-esb-test/internal/config"
	"go-esb-test/internal/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	SeedUsersAndItems(db)
	log.Println("Database seeding completed successfully.")
}

func mustHashPassword(password string) string {
	hash, err := hashPassword(password)
	if err != nil {
		log.Fatalf("Failed to hash password: %+v", err)
	}
	return hash
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func SeedUsersAndItems(db *gorm.DB) {
	// Create example users
	users := []entity.User{
		{
			Fullname:       "John Doe",
			Username:       "johndoe",
			Email:          "john.doe@example.com",
			Phone:          "0800000001",
			Address:        "123 Main St",
			HashedPassword: mustHashPassword("password123"),
		},
		{
			Fullname:       "Jane Smith",
			Username:       "janesmith",
			Email:          "jane.smith@example.com",
			Phone:          "0800000002",
			Address:        "456 Elm St",
			HashedPassword: mustHashPassword("secret123"),
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to seed user data: %+v", err)
		}
	}

	// Create example items
	items := []entity.Item{
		{
			Name:        "Item A",
			Type:        "Type A",
			Description: "Description for Item A",
		},
		{
			Name:        "Item B",
			Type:        "Type B",
			Description: "Description for Item B",
		},
	}

	for _, item := range items {
		if err := db.Create(&item).Error; err != nil {
			log.Fatalf("Failed to seed item data: %+v", err)
		}
	}
}
