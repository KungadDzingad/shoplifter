package main

import (
	"fmt"
	"os"

	"github.com/KungadDzingad/shoplifter-gateway/src/messaging"
	"github.com/KungadDzingad/shoplifter-gateway/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	messaging.InitKafkaConsumer()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("GATEWAY_PORT")))
}

func GetCoreURL() string {
	return fmt.Sprintf("http://%s:%s", os.Getenv("CORE_HOST"), os.Getenv("CORE_PORT"))
}
