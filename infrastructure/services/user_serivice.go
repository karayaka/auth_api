package services

import (
	"auth_api/core/common"
	customerrors "auth_api/core/custom_errors"
	models "auth_api/core/models/base_models"
	user_model "auth_api/core/models/user_models"
	usercachemodels "auth_api/infrastructure/dto_models/user_cache_models"
	userdtomodel "auth_api/infrastructure/dto_models/user_dto_model"
	rmqproviders "auth_api/infrastructure/providers/rmq_providers"
	repositorys "auth_api/infrastructure/repositorys"
	"auth_api/presentation/user_api/view_models/request"
	"fmt"
)

type IUserService interface {
	AddUser(model userdtomodel.UserCreateDto) error
	UpdateUser(model userdtomodel.UserUpdateDto) error
	DeleteUser(id uint) error
	GetUserByEmail(email string) (*userdtomodel.UserDto, error)
	GetAllUser() ([]userdtomodel.UserDto, error)
	Login(model request.LoginRequestModel) (string, error)
}

type UserService struct {
	uow  *repositorys.UnitOfWork
	rmqp *rmqproviders.RmqProvider
}

func NewUserService(_uow *repositorys.UnitOfWork, rmqp *rmqproviders.RmqProvider) IUserService {
	return &UserService{
		uow:  _uow,
		rmqp: rmqp,
	}
}

func (us UserService) AddUser(model userdtomodel.UserCreateDto) error {
	if model.Password == "" {
		return customerrors.NewNotFoundError("Pasaport bilgisi boş olamaz")
	}
	hashedPasword, err := common.HashPassword(model.Password)
	if err != nil {
		return err
	}
	us.rmqp.UserRmqProvider.AddMeesageToEvent(model.Name)
	redisErr := us.uow.UserCacheRepository.AddUser(usercachemodels.UserCacheModel{
		ID:      model.ID,
		Name:    model.Name,
		Email:   model.Email,
		Surname: model.Surname,
	})
	if redisErr != nil {
		fmt.Println(redisErr.Error())
	}

	us.uow.UserRepository.Add(&user_model.UserEntity{
		Name:     model.Name,
		Surname:  model.Surname,
		Email:    model.Email,
		Password: hashedPasword,
		BaseEntitiy: models.BaseEntitiy{
			CreatedBy: model.SessionId,
			UpdatedBy: model.SessionId,
		},
	})
	return nil
}
func (us UserService) UpdateUser(model userdtomodel.UserUpdateDto) error {
	if model.Password == "" {
		//hata yönetimine bakılacak
	}
	user, err := us.uow.UserRepository.GetByID(model.ID)
	if err != nil {
		return customerrors.NewNotFoundError("Kullanıcı Bulunamadı")
	}
	user.Name = model.Name
	user.Surname = model.Surname
	user.Password = model.Password
	user.Email = model.Email
	us.uow.UserRepository.Update(user)
	return nil
}

func (us UserService) DeleteUser(id uint) error {
	err := us.uow.UserRepository.Delete(id)
	if err != nil {
		return customerrors.NewCustomError("Silme işleminde bir hata oluştu")
	}
	return nil
}

func (us UserService) GetUserByEmail(email string) (*userdtomodel.UserDto, error) {
	user, err := us.uow.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, customerrors.NewNotFoundError("Kullanıcı Bulunamadı")
	}
	return &userdtomodel.UserDto{
		ID:      user.ID,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
	}, nil
}

func (us UserService) GetAllUser() ([]userdtomodel.UserDto, error) {
	users, err := us.uow.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	var retVals []userdtomodel.UserDto
	for _, user := range users {

		retVals = append(retVals, userdtomodel.UserDto{
			ID:      user.ID,
			Name:    user.Name,
			Surname: user.Surname,
			Email:   user.Email,
		})
	}
	return retVals, nil
}

func (us UserService) Login(model request.LoginRequestModel) (string, error) {

	user, err := us.uow.UserRepository.GetByEmail(model.Email)
	if err != nil || user == nil {
		return "", customerrors.NewUnAuthorizedError("Kullancı Adı veya Şifre Hatalı")
	}
	if !common.VerifyPassword(model.Password, user.Password) {
		return "", customerrors.NewUnAuthorizedError("Kullancı Adı veya Şifre Hatalı")
	}
	token, tokenErr := common.GenerateJWT(*user)

	return *token, tokenErr
}
