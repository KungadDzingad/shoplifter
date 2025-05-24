package handlers

import (
	"fmt"

	"github.com/KungadDzingad/shoplifter-common/models"
	database "github.com/KungadDzingad/shoplifter-core/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func PostUser(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		log.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	defer log.Info(fmt.Sprintf("User %s saved", user.Username))

	database.CONNECTION.Db.Create(&user)

	return ctx.Status(200).JSON(user)
}
