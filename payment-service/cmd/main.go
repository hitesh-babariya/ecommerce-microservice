package main

import (
	"ecommerce-microservice/payment-service/internal/handler"
	"ecommerce-microservice/payment-service/internal/infrastructure"
	"ecommerce-microservice/payment-service/internal/repository"
	"ecommerce-microservice/payment-service/internal/usecase"
	"log"
)

func main() {

	// -------------------------
	// Database
	// -------------------------

	db, err := infrastructure.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// -------------------------
	// Repository
	// -------------------------

	paymentRepo :=
		repository.NewSQLitePaymentRepository(db)

	// -------------------------
	// Payment Processor
	// -------------------------

	paymentProcessor :=
		usecase.NewMockPaymentProcessor()

	// -------------------------
	// Kafka Producer
	// -------------------------

	paymentPublisher :=
		infrastructure.NewKafkaProducer(
			[]string{
				"localhost:9092",
			},
		)

	defer paymentPublisher.Close()

	// -------------------------
	// Usecase
	// -------------------------

	paymentUsecase :=
		usecase.NewPaymentUsecase(
			paymentRepo,
			paymentProcessor,
			paymentPublisher,
		)

	// -------------------------
	// Handler
	// -------------------------

	paymentHandler :=
		handler.NewPaymentHandler(
			paymentUsecase,
		)

	// -------------------------
	// Router
	// -------------------------

	router :=
		infrastructure.NewRouter(
			paymentHandler,
		)

	// -------------------------
	// Server
	// -------------------------

	log.Println(
		"Payment Service running on :8082",
	)

	if err := router.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
