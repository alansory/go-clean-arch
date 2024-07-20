package repository

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T, id int64) error {
	return db.Where("id = ?", id).Delete(entity).Error
}

func (r *Repository[T]) CountById(db *gorm.DB, id int64) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (r *Repository[T]) FindById(db *gorm.DB, entity *T, id int64, preloads ...string) error {
	query := db.Where("id = ?", id)

	// Apply preloads if provided
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	return query.Take(entity).Error
}
