package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
)

// /deneme amaçlı midleware
func AuthMiddlewareDeneme(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-API-KEY")

		// Anahtar kontrolü
		if apiKey != "secret123" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}

		// Eğer anahtar doğruysa, isteği bir sonraki handler'a iletir
		return next(c)
	}
}
