package infrastructure

import (
	"context"
	"ecommerce-microservice/order-service/internal/usecase"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type PaymentEvent struct {
	PaymentID int64  `json:"paymentId"`
	OrderID   int64  `json:"orderId"`
	UserID    int64  `json:"userId"`
	Amount    int64  `json:"amount"`
	Status    string `json:"status"`
}

type PaymentConsumer struct {
	reader       *kafka.Reader
	orderUsecase *usecase.OrderUsecase
}

func NewPaymentConsumer(
	topic string,
	groupID string,
) *PaymentConsumer {

	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: []string{
				"localhost:9092",
			},
			Topic:   topic,
			GroupID: groupID,
		},
	)

	return &PaymentConsumer{
		reader: reader,
	}
}

func (c *PaymentConsumer) Start(
	ctx context.Context,
) {

	for {

		message, err :=
			c.reader.ReadMessage(ctx)

		if err != nil {

			if ctx.Err() != nil {
				return
			}

			log.Println(
				"Kafka read error:",
				err,
			)

			continue
		}

		var event PaymentEvent

		if err := json.Unmarshal(
			message.Value,
			&event,
		); err != nil {

			log.Println(
				"Invalid payment event:",
				err,
			)

			continue
		}

		log.Printf(
			"Payment event received: paymentID=%d orderID=%d status=%s\n",
			event.PaymentID,
			event.OrderID,
			event.Status,
		)

	}
}

func (c *PaymentConsumer) Close() error {
	return c.reader.Close()
}
