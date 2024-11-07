package databaserepositorys

import "gorm.io/gorm"

type IBaseRepository[T any] interface {
	Add(model *T) error
	Update(model *T) error
	Delete(id uint) error
	GetByID(id uint) (*T, error)
	GetAll() ([]T, error)
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

// NewBaseRepository, yeni bir BaseRepository nesnesi oluşturur.
func NewBaseRepository[T any](db *gorm.DB) IBaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

// GetByID, belirtilen ID'ye göre kaydı döndürür.
func (r *BaseRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	result := r.db.First(&entity, id)
	return &entity, result.Error
}

// GetAll, tüm kayıtları döndürür.
func (r *BaseRepository[T]) GetAll() ([]T, error) {
	var entities []T
	result := r.db.Find(&entities)
	return entities, result.Error
}

// Create, yeni bir kayıt oluşturur.
func (r *BaseRepository[T]) Add(entity *T) error {
	result := r.db.Create(&entity)
	return result.Error
}

// Update, mevcut bir kaydı günceller.
func (r *BaseRepository[T]) Update(entity *T) error {
	result := r.db.Save(&entity)
	return result.Error
}

// Delete, belirtilen ID'ye göre kaydı siler.
func (r *BaseRepository[T]) Delete(id uint) error {
	var entity T
	result := r.db.Delete(&entity, id)
	return result.Error
}
