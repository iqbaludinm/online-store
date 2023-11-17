package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	app := fiber.New()

	port := os.Getenv("APP_PORT")
	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
