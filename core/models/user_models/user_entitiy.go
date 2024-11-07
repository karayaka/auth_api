package user_model

import (
	models "auth_api/core/models/base_models"
)

type UserEntity struct {
	models.BaseEntitiy
	Name     string
	Surname  string
	Email    string
	Password string
}

func (rec UserEntity) TableName() string {
	return "users"
}
