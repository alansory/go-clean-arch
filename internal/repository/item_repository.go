package repository

import (
	"go-esb-test/internal/entity"
	"go-esb-test/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ItemRepository struct {
	Repository[entity.Item]
	Log *logrus.Logger
}

func NewItemRepository(log *logrus.Logger) *ItemRepository {
	return &ItemRepository{
		Log: log,
	}
}

func (r *ItemRepository) Search(db *gorm.DB, request *model.SearchItemRequest) ([]entity.Item, int64, error) {
	var items []entity.Item

	query := db.Scopes(r.FilterItem(request)).
		Offset((request.Page - 1) * request.PerPage).
		Limit(request.PerPage)

	if err := query.Find(&items).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	totalQuery := db.Model(&entity.Item{}).
		Scopes(r.FilterItem(request))

	if err := totalQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *ItemRepository) FilterItem(request *model.SearchItemRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {

		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}

		if description := request.Description; description != "" {
			description = "%" + description + "%"
			tx = tx.Where("description LIKE ?", description)
		}

		return tx
	}
}
