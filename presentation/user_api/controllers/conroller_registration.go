package controllers

import (
	customerrors "auth_api/core/custom_errors"
	rmqproviders "auth_api/infrastructure/providers/rmq_providers"
	repositorys "auth_api/infrastructure/repositorys"
	"auth_api/infrastructure/services"
	"auth_api/presentation/user_api/view_models/response"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func RegisterControllers(e *echo.Echo, uow *repositorys.UnitOfWork, rmqp *rmqproviders.RmqProvider) {

	uc := NewUserConroller(services.NewUserService(uow, rmqp))
	uc.RegisterUserController(e)
	sc := NewSecurityController(services.NewUserService(uow, rmqp))
	sc.RegisterSecurityController(e)
	//diğer kotrollerde bursa eklecek
}

func Ok[T any](c echo.Context, result T) error {
	res := response.BaseResponse[T]{
		Data:    result,
		Message: "İşlem Başarılı",
		Date:    time.Now(),
	}
	c.JSON(http.StatusOK, res)
	return nil
}
func ErrorResponse(c echo.Context, err error) error {
	res := response.ErrorResponse{
		Message: err.Error(),
		Date:    time.Now(),
	}
	var customError *customerrors.CustomError
	var notFoundError *customerrors.NotFoundError
	var unauthorizederror *customerrors.UnAuthorizedError
	if errors.As(err, &customError) {
		c.JSON(http.StatusBadRequest, res)
	} else if errors.As(err, &notFoundError) {
		c.JSON(http.StatusNotFound, res)
	} else if errors.As(err, &unauthorizederror) {
		c.JSON(http.StatusUnauthorized, res)
	} else {
		c.JSON(http.StatusBadRequest, res)
	}
	return nil
}
