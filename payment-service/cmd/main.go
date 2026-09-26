package main

import (
	"context"
	"ecommerce-microservice/payment-service/internal/handler"
	"ecommerce-microservice/payment-service/internal/infrastructure"
	"ecommerce-microservice/payment-service/internal/repository"
	"ecommerce-microservice/payment-service/internal/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	// HTTP Server
	// -------------------------

	server := &http.Server{
		Addr:    ":8083",
		Handler: router,
	}

	go func() {

		log.Println(
			"Payment Service running on :8083",
		)

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			log.Fatalf(
				"server error: %v",
				err,
			)
		}

	}()

	// -------------------------
	// Graceful Shutdown
	// -------------------------

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	log.Println(
		"Shutdown signal received...",
	)

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {

		log.Printf(
			"Server forced to shutdown: %v",
			err,
		)
	}

	log.Println(
		"Payment service stopped",
	)

}
