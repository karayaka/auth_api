package controllers

import (
	"auth_api/infrastructure/services"
	"auth_api/presentation/user_api/middlewares"
	"auth_api/presentation/user_api/view_models/request"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type UserConroller struct {
	UserService services.IUserService
}

func NewUserConroller(us services.IUserService) *UserConroller {
	return &UserConroller{
		UserService: us,
	}
}

func (uc UserConroller) RegisterUserController(e *echo.Echo) {
	//e.GET("api/user", uc.GetUsers, middlewares.AuthMiddleware)
	g := e.Group("api/user", middlewares.AuthMiddleware) //groplayıp koruma sağlanıyor
	g.GET("/GetAllUser", uc.GetUsers)
	g.POST("/Add", uc.AddUser)
	//e.POST("api/login", uc.Login)
}

func (uc UserConroller) GetUsers(c echo.Context) error {
	users, err := uc.UserService.GetAllUser()
	if err != nil {
		return ErrorResponse(c, err)
	}
	return Ok(c, users)
}
func (uc UserConroller) AddUser(c echo.Context) error {
	var user request.UserCreateRequestModel
	sessionId, _ := strconv.ParseUint(c.Get("sesionId").(string), 10, 32)

	err := c.Bind(&user)
	if err != nil {
		return ErrorResponse(c, err)
	}
	///todo validation yapısı daha detaylı çalışılacak mesajların özelleştirilmesi mdel yapısı vs
	validate := validator.New(validator.WithRequiredStructEnabled())
	valerr := validate.Struct(&user)
	if valerr != nil {
		if validationErrors, ok := valerr.(validator.ValidationErrors); ok {
			return ErrorResponse(c, validationErrors)
		} else {
			// Başka bir hata türü
		}
	}

	addErr := uc.UserService.AddUser(user.ToUserCreateModel(uint(sessionId)))
	if addErr != nil {
		return ErrorResponse(c, addErr)
	}
	return Ok(c, "")
}
