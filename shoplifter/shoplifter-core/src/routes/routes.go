package routes

import (
	"github.com/KungadDzingad/shoplifter-core/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/user", handlers.PostUser)
}
