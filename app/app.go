package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/iqbaludinm/online-store/config"
	"github.com/iqbaludinm/online-store/handlers"
	"github.com/iqbaludinm/online-store/helpers"
	"github.com/iqbaludinm/online-store/repositories"
	routesV1 "github.com/iqbaludinm/online-store/routes"
	"github.com/iqbaludinm/online-store/services"
)

var app = fiber.New()

func StartApp() {
	repo := repositories.NewRepo(config.GORM.DB)
	services := services.NewService(repo)
	server := handlers.NewHttpServer(services)

	app.Use(logger.New())
	app.Use(cors.New())

	routesV1.RegisterAPIs(app, server)
	
	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return helpers.NotFound(c, "Route not found")
	})
	

	port := os.Getenv("APP_PORT")
	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
