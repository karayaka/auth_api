package request

import (
	basemodels "auth_api/infrastructure/dto_models/base_models"
	userdtomodel "auth_api/infrastructure/dto_models/user_dto_model"
)

type UserCreateRequestModel struct {
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" valdate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

func (ucrm UserCreateRequestModel) ToUserCreateModel(sessionId uint) userdtomodel.UserCreateDto {
	return userdtomodel.UserCreateDto{
		UserDto: userdtomodel.UserDto{
			BaseDTO: basemodels.BaseDTO{
				SessionId: sessionId,
			},
			Name:    ucrm.Name,
			Surname: ucrm.Surname,
			Email:   ucrm.Email,
		},
		Password: ucrm.Password,
	}
}
