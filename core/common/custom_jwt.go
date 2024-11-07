package common

import (
	customerrors "auth_api/core/custom_errors"
	user_model "auth_api/core/models/user_models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

type CustomJWT struct {
	Name    string
	Surname string
	Email   string
	jwt.RegisteredClaims
}

func GenerateJWT(user user_model.UserEntity) (*string, error) {
	userClaims := CustomJWT{
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.FormatUint(uint64(user.ID), 10),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	accesToken := jwt.NewWithClaims(jwt.SigningMethodHS512, userClaims)
	signedAccesToken, err := accesToken.SignedString([]byte(os.Getenv("ACCESS_SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &signedAccesToken, nil
}

func GetSession(c echo.Context) (*CustomJWT, error) {
	c.Response().Header().Add("Vary", "Authorization")
	authHeader := c.Request().Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer") {
		return nil, customerrors.NewUnAuthorizedError("Hatalı Token")
	}
	tokens := strings.Split(authHeader, " ")
	if len(tokens) < 2 {
		return nil, customerrors.NewUnAuthorizedError("Hatalı Token")
	}
	return ParseJWT(tokens[1])
}

func ParseJWT(accesToken string) (*CustomJWT, error) {

	parsedAccesToken, err := jwt.ParseWithClaims(accesToken, &CustomJWT{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
	}, jwt.WithLeeway(5*time.Second))
	if !parsedAccesToken.Valid {
		return nil, customerrors.NewUnAuthorizedError("Token Hatası")
	}
	if err != nil {
		return nil, err
	} else if claims, ok := parsedAccesToken.Claims.(*CustomJWT); ok {
		return claims, nil
	} else {
		return nil, customerrors.NewUnAuthorizedError("Token Hatası")
	}
}
