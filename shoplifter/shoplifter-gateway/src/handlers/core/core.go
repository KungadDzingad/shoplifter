package handlers_core

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/KungadDzingad/shoplifter-gateway/src/messaging"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func getKafkaWriter() *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "core-requests",
		Balancer: &kafka.LeastBytes{},
	})
}

func GetUsers(c *fiber.Ctx) error {
	correlationID := uuid.New().String()

	reqMsg := map[string]string{
		"correlation_id": correlationID,
		"request_type":   "get-users",
	}
	msgBytes, _ := json.Marshal(reqMsg)

	writer := getKafkaWriter()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(correlationID),
		Value: msgBytes,
	})

	if err != nil {
		log.Error("Failed to write Kafka message:", err)
		return c.Status(500).SendString("Internal Error")
	}

	respChan := make(chan messaging.ResponseEnvelope)
	messaging.Lock.Lock()
	messaging.ResponseChannels[correlationID] = respChan
	messaging.Lock.Unlock()

	select {
	case resp := <-respChan:
		log.Info("Elo")
		return c.Status(resp.StatusCode).Send(resp.Payload)
	case <-time.After(5 * time.Second):
		return c.Status(504).SendString("Gateway Timeout")
	}
}

func PostUser(c *fiber.Ctx) error {
	type UserCreation struct {
		Mail     string "json:mail"
		Username string "json:username"
		Password string "json:password"
	}

	user := new(UserCreation)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	jsonData, _ := json.Marshal(user)

	resp, err := http.Post(messaging.GetCoreUrl()+"/user", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to connect to core service",
		})
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return c.Status(resp.StatusCode).Send(body)
}
