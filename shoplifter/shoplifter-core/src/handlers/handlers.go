package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
)

type KafkaRequest struct {
	CorrelationID string `json:"correlation_id"`
	RequestType   string `json:"request_type"`
}

type KafkaResponse struct {
	CorrelationID string          `json:"correlation_id"`
	StatusCode    int             `json:"status_code"`
	Payload       json.RawMessage `json:"payload"`
}

func getKafkaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "core-requests",
		GroupID: "core-service",
	})
}

func getKafkaWriter() *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "gateway-responses",
	})
}

func ListenForKafkaRequests() {
	reader := getKafkaReader()
	writer := getKafkaWriter()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}

		var req KafkaRequest
		if err := json.Unmarshal(message.Value, &req); err != nil {
			log.Println("Invalid request:", err)
			continue
		}

		switch req.RequestType {
		case "get_users":
			handleGetUsers(req.CorrelationID, writer)
		default:
			log.Println("Unknown request_type:", req.RequestType)
		}
	}
}

func Home(c *fiber.Ctx) error {
	return c.SendString("eluwina")
}

func sendKafkaResponse(writer *kafka.Writer, correlationID string, status int, payload []byte) {
	resp := KafkaResponse{
		CorrelationID: correlationID,
		StatusCode:    status,
		Payload:       payload,
	}
	msg, err := json.Marshal(resp)
	if err != nil {
		log.Println("Kafka response marshal error:", err)
		return
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(correlationID),
		Value: msg,
	})
	if err != nil {
		log.Println("Kafka write error:", err)
	}
}
