package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

var coreUrl string

var ResponseChannels = make(map[string]chan ResponseEnvelope)
var Lock sync.Mutex

type KafkaResponse struct {
	CorrelationID string          `json:"correlation_id"`
	StatusCode    int             `json:"status_code"`
	Payload       json.RawMessage `json:"payload"`
}

type ResponseEnvelope struct {
	StatusCode int
	Payload    []byte
}

func GetCoreUrl() string {
	if &coreUrl == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		coreUrl = fmt.Sprintf("http://%s:%s", os.Getenv("CORE_HOST"), os.Getenv("CORE_PORT"))
	}
	return coreUrl
}

func listenForKafkaResponse() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "gateway-responses",
		GroupID: "gateway",
	})

	go func() {
		for {
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Println("Error reading kafka response", err)
				continue
			}
			var msg KafkaResponse
			if err := json.Unmarshal(m.Value, &msg); err != nil {
				log.Println("Invalid Kafka message:", err)
				continue
			}

			Lock.Lock()
			if ch, ok := ResponseChannels[msg.CorrelationID]; ok {
				ch <- ResponseEnvelope{
					StatusCode: msg.StatusCode,
					Payload:    msg.Payload,
				}
				close(ch)
				delete(ResponseChannels, msg.CorrelationID)
			}
			Lock.Unlock()
		}
	}()

}

func InitKafkaConsumer() {
	listenForKafkaResponse()
}
