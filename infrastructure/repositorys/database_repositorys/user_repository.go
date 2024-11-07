package databaserepositorys

import (
	user_models "auth_api/core/models/user_models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	IBaseRepository[user_models.UserEntity]
	GetByEmail(email string) (*user_models.UserEntity, error)
}

type UserRepository struct {
	IBaseRepository[user_models.UserEntity]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		IBaseRepository: NewBaseRepository[user_models.UserEntity](db),
		db:              db,
	}
}

func (r *UserRepository) GetByEmail(email string) (*user_models.UserEntity, error) {
	var user user_models.UserEntity
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}
