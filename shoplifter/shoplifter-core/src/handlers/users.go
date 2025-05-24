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

	database.DB.Db.Create(&user)
	return ctx.Status(200).JSON(user)
}

func GetUsers(ctx *fiber.Ctx) error {
	users := []models.User{}

	database.DB.Db.Find(&users)
	return ctx.Status(200).JSON(users)
}
