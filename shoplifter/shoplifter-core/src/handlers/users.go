package handlers

import (
	"encoding/json"

	"github.com/KungadDzingad/shoplifter-common/models"
	database "github.com/KungadDzingad/shoplifter-core/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/segmentio/kafka-go"
)

func PostUser(c *fiber.Ctx) {
	// user := new(models.User)
	// db := database.Connection()

	// if db == nil {
	// 	return c.SendString("Database issue")
	// }

	// if err := c.BodyParser(user); err != nil {
	// 	log.Error(err.Error())
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": err.Error(),
	// 	})
	// }

	// defer log.Info(fmt.Sprintf("User %s saved", user.Username))

	// database.Connection().Db.Create(&user)
	// return c.Status(200).JSON(user)
}

func handleGetUsers(correlationID string, writer *kafka.Writer) {
	db := database.Connection()
	if db == nil {
		sendKafkaResponse(writer, correlationID, 500, []byte(`"Database issue"`))
		return
	}

	var users []models.User
	if err := db.Db.Find(&users).Error; err != nil {
		sendKafkaResponse(writer, correlationID, 500, []byte(`"Database error"`))
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		sendKafkaResponse(writer, correlationID, 500, []byte(`"Serialization error"`))
		return
	}
	log.Info("get-users", data)

	sendKafkaResponse(writer, correlationID, 200, data)
}
