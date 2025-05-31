package routes_core

import (
	handlers_core "github.com/KungadDzingad/shoplifter-gateway/src/handlers/core"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/user", handlers_core.GetUsers)
	// app.Post("/user", handlers_core.PostUser)
}
