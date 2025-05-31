package routes

import (
	routes_core "github.com/KungadDzingad/shoplifter-gateway/src/routes/core"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	routes_core.SetupRoutes(app)
}
