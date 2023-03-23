package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		AppName:       "GO APP",
		CaseSensitive: true,
		ServerHeader:  "GO APP",
	})
	configs.ConnectDB()

	routes.UserRoute(app)
	app.Static("/", "./public")
	app.Listen(":3000")
}
