package main

import (
	database "github.com/KungadDzingad/shoplifter-core/src/database"
	"github.com/KungadDzingad/shoplifter-core/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
