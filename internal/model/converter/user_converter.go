package converter

import (
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
