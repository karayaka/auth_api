package response

import userdtomodel "auth_api/infrastructure/dto_models/user_dto_model"

type UserResponseModel struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

func FromUserDtoMdel(model userdtomodel.UserDto) *UserResponseModel {
	return &UserResponseModel{
		Name:    model.Name,
		Surname: model.Surname,
		Email:   model.Email,
		ID:      model.ID,
	}
}
