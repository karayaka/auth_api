package controllers

import (
	"auth_api/infrastructure/services"
	"auth_api/presentation/user_api/view_models/request"

	"github.com/labstack/echo"
)

type SecurityController struct {
	UserService services.IUserService
}

func (sc SecurityController) RegisterSecurityController(e *echo.Echo) {
	e.POST("api/security/login", sc.Login)
}

func NewSecurityController(us services.IUserService) *SecurityController {
	return &SecurityController{
		UserService: us,
	}
}

func (uc SecurityController) Login(c echo.Context) error {
	var user request.LoginRequestModel
	err := c.Bind(&user)
	if err != nil {
		return ErrorResponse(c, err)
	}
	jwt, loginErr := uc.UserService.Login(user)
	if loginErr != nil {
		return ErrorResponse(c, loginErr)
	}
	return Ok(c, jwt)

}
