package userdtomodel

import basemodels "auth_api/infrastructure/dto_models/base_models"

type UserDto struct {
	basemodels.BaseDTO
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}
