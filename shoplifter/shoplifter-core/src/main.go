package main

import (
	"fmt"
	"os"

	database "github.com/KungadDzingad/shoplifter-core/src/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Listen(fmt.Sprintf("0.0.0.0:%s", os.Getenv("CORE_PORT")))
}
