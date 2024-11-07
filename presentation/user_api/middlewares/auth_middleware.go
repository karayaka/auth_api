package middlewares

import (
	"auth_api/core/common"
	"auth_api/presentation/user_api/view_models/response"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Vary", "Authorization")
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer") {
			return unauthorizedResponse(c, "Giriş bilgileri hatalı")
		}
		tokens := strings.Split(authHeader, " ")
		if len(tokens) < 2 {
			return unauthorizedResponse(c, "Giriş bilgileri hatalı")
		}
		user, err := common.ParseJWT(tokens[1])
		if err != nil {
			return unauthorizedResponse(c, "Giriş bilgileri hatalı")
		}
		c.Set("sesionId", user.ID)

		//return unauthorizedResponse(c, user.Name)
		return next(c)
	}
}

func unauthorizedResponse(c echo.Context, msg string) error {
	res := response.ErrorResponse{
		Message: msg,
		Date:    time.Now(),
	}
	return c.JSON(http.StatusUnauthorized, res)
}
