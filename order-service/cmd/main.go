package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ecommerce-microservice/order-service/internal/handler"
	framework "ecommerce-microservice/order-service/internal/infrastructure"
	"ecommerce-microservice/order-service/internal/model"
	"ecommerce-microservice/order-service/internal/repository"
	"ecommerce-microservice/order-service/internal/usecase"
)

func main() {

	// ==========================
	// Infrastructure
	// ==========================

	db, err := framework.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// ==========================
	// Repository
	// ==========================

	orderRepo := repository.NewSQLiteOrderRepository(db)

	// ==========================
	// User Service HTTP Client
	// ==========================

	userClient := framework.NewHTTPUserServiceClient(
		"http://localhost:8081",
	)

	// ==========================
	// Usecase
	// ==========================

	orderUsecase := usecase.NewOrderUsecase(orderRepo, userClient)

	// ==========================
	// Payment Consumers
	// ==========================

	// paymentConsumer := framework.NewPaymentConsumer()

	paymentSuccessConsumer := framework.NewPaymentConsumer(
		"payment-success",
		"order-service-success",
		model.OrderStatusPaid,
		orderUsecase,
	)

	paymentFailedConsumer := framework.NewPaymentConsumer(
		"payment-failed",
		"order-service-failed",
		model.OrderStatusCancelled,
		orderUsecase,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go paymentSuccessConsumer.Start(ctx)
	defer paymentFailedConsumer.Close()

	// ==========================
	// Handler
	// ==========================

	orderHandler := handler.NewOrderHandler(orderUsecase)

	// ==========================
	// Router
	// ==========================

	router := framework.NewRouter(orderHandler)

	// ==========================
	// HTTP Server
	// ==========================

	server := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	go func() {

		log.Println("Order service running on :8082")

		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {

			log.Fatalf("server error: %v", err)
		}

	}()

	// ==========================
	// Graceful Shutdown
	// ==========================

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	log.Println("Shutdown signal received...")

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

	log.Println("Order service stopped")
}
