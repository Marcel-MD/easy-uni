package data

import (
	"gorm.io/gorm"
)

type Repository[T any] interface {
	FindAll(page, size int) ([]T, error)
	FindById(id string) (T, error)
	Create(t *T) error
	Update(t *T) error
	Delete(t *T) error
}

func NewRepository[T any]() Repository[T] {
	return &repository[T]{
		db: GetDB(),
	}
}

type repository[T any] struct {
	db *gorm.DB
}

func (r *repository[T]) FindAll(page, size int) ([]T, error) {
	var ts []T

	err := r.db.Scopes(paginate(page, size)).Find(&ts).Error
	return ts, err
}

func paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 50
		}

		if page <= 0 {
			page = 1
		}

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

func (r *repository[T]) FindById(id string) (T, error) {
	var t T
	err := r.db.First(&t, "id = ?", id).Error
	return t, err
}

func (r *repository[T]) Create(t *T) error {
	return r.db.Create(t).Error
}

func (r *repository[T]) Update(t *T) error {
	return r.db.Save(t).Error
}

func (r *repository[T]) Delete(t *T) error {
	return r.db.Delete(t).Error
}
