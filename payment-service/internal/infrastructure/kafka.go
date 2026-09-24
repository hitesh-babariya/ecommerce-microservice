package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"

	"ecommerce-microservice/payment-service/internal/model"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(
	brokers []string,
) *KafkaProducer {

	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{
		writer: writer,
	}
}

type PaymentEvent struct {
	PaymentID int64   `json:"paymentId"`
	OrderID   int64   `json:"orderId"`
	UserID    int64   `json:"userId"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
}

func (p *KafkaProducer) publish(
	ctx context.Context,
	topic string,
	payment model.Payment,
) error {

	event := PaymentEvent{
		PaymentID: payment.ID,
		OrderID:   payment.OrderID,
		UserID:    payment.UserID,
		Amount:    payment.Amount,
		Status:    payment.Status,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf(
			"marshal payment event: %w",
			err,
		)
	}

	message := kafka.Message{
		Topic: topic,
		Key: []byte(
			fmt.Sprintf("%d", payment.OrderID),
		),
		Value: data,
		Time:  time.Now(),
	}

	if err := p.writer.WriteMessages(
		ctx,
		message,
	); err != nil {
		return fmt.Errorf(
			"publish payment event: %w",
			err,
		)
	}

	return nil
}

func (p *KafkaProducer) PublishPaymentSuccess(
	ctx context.Context,
	payment model.Payment,
) error {

	return p.publish(
		ctx,
		"payment-success",
		payment,
	)
}

func (p *KafkaProducer) PublishPaymentFailed(
	ctx context.Context,
	payment model.Payment,
) error {

	return p.publish(
		ctx,
		"payment-failed",
		payment,
	)
}

func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}
