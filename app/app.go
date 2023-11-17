package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/iqbaludinm/online-store/config"
	"github.com/iqbaludinm/online-store/handlers"
	"github.com/iqbaludinm/online-store/repositories"
	"github.com/iqbaludinm/online-store/services"
)

var app = fiber.New()

func StartApp() {
	repo := repositories.NewRepo(config.GORM.DB)
	services := services.NewService(repo)
	_ = handlers.NewHttpServer(services)

	port := os.Getenv("APP_PORT")
	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
