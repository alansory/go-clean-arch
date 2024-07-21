package usecase

import (
	"context"
	"go-esb-test/internal/model"
	"go-esb-test/internal/model/converter"
	"go-esb-test/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserUseCase(
	db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	UserRepository *repository.UserRepository,
) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		UserRepository: UserRepository,
	}
}

func (c *UserUseCase) List(ctx context.Context, request *model.SearchUserRequest) ([]model.UserResponse, int64, error) {
	users, total, err := c.UserRepository.Search(c.DB, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting users")
		return nil, 0, err
	}

	responses := make([]model.UserResponse, len(users))
	for i, user := range users {
		responses[i] = *converter.UserToResponse(&user)
	}

	return responses, total, nil
}
