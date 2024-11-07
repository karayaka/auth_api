package main

import (
	"auth_api/infrastructure/providers/consumers"
	rmqproviders "auth_api/infrastructure/providers/rmq_providers"
	repositorys "auth_api/infrastructure/repositorys"
	"auth_api/persistence"
	"auth_api/presentation/user_api/controllers"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
)

func main() {
	err := godotenv.Load()
	e := echo.New()
	if err != nil {
		e.Logger.Fatal("Error loading .env file")
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	port := os.Getenv("APP_PORT")
	connectionString := os.Getenv("USER_DB_CONNECTION")
	context := persistence.RegisterUserContext(connectionString, e.Logger)
	db := context.Init()
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("USER_CACHE_DB_CONNECTION"),
	})

	// Kapanış işlemi için bir kanal oluşturma
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Asenkron olarak uygulama kapanışını dinleyin
	go func() {
		<-quit
		e.Logger.Fatal("Kapanış işlemi başlatılıyor...")
		if err := redisClient.Close(); err != nil {
			e.Logger.Fatal("Redis kapanırken hata oluştu: %v", err)
		}
		e.Logger.Fatal("Redis bağlantısı kapatıldı")
		os.Exit(0)
	}()

	rmqUri := os.Getenv("RMQ_CONNECTION")
	connection, rmqErr := amqp.Dial(rmqUri)
	if rmqErr != nil {
		e.Logger.Fatal("RMQ Başlatılamadı")
	}
	channel, channelErr := connection.Channel()
	if channelErr != nil {
		e.Logger.Fatal("RMQ Kapanış işlemi başlatılıyor...")
	}
	go func() {
		<-quit
		e.Logger.Fatal("RMQ Kapanış işlemi başlatılıyor...")
		if err := connection.Close(); err != nil {
			e.Logger.Fatal("RMQ kapanırken hata oluştu: %v", err)
		}
		channel.Close()
		e.Logger.Fatal("RMQ bağlantısı kapatıldı")
		os.Exit(0)
	}()
	prv := rmqproviders.NewRmqProvider(channel)

	consumers.RegisterConsumer(channel)
	//veir tabanı erişim katmanı oluşturuldu
	uow := repositorys.NewUnitOfWork(redisClient, db)
	//conroller register edildi
	controllers.RegisterControllers(e, uow, prv)
	//e.Use(middlewares.AuthMiddleware)

	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))

}
