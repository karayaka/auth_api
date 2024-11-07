package persistence

import (
	user_model "auth_api/core/models/user_models"
	"log"

	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserContext struct {
	ConnectionString string
	log              *echo.Logger
}

func RegisterUserContext(cs string, log echo.Logger) *UserContext {
	return &UserContext{
		ConnectionString: cs,
		log:              &log,
	}
}

func (uc UserContext) Init() *gorm.DB {

	db, err := gorm.Open(postgres.Open(uc.ConnectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Hata Oluştu")
	}
	log.Default().Println("deneme")
	db.AutoMigrate(user_model.UserEntity{})

	return db
}
