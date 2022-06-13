package main

import (
	"backend/configs"
	"backend/routes" //add this

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	app.Use(cors.New())

	//routes
	routes.UserRoute(app) //add this

	app.Listen(":6001")
}
