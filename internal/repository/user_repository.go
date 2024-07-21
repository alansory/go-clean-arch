package repository

import (
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) Search(db *gorm.DB, request *model.SearchUserRequest) ([]entity.User, int64, error) {
	var users []entity.User

	query := db.Scopes(r.FilterItem(request)).
		Offset((request.Page - 1) * request.PerPage).
		Limit(request.PerPage)

	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	totalQuery := db.Model(&entity.User{}).
		Scopes(r.FilterItem(request))

	if err := totalQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *UserRepository) FilterItem(request *model.SearchUserRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {

		if fullname := request.Fullname; fullname != "" {
			fullname = "%" + fullname + "%"
			tx = tx.Where("fullname LIKE ?", fullname)
		}

		if email := request.Email; email != "" {
			email = "%" + email + "%"
			tx = tx.Where("email LIKE ?", email)
		}

		if username := request.Username; username != "" {
			username = "%" + username + "%"
			tx = tx.Where("username LIKE ?", username)
		}

		return tx
	}
}
