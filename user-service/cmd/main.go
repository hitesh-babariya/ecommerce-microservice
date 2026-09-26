package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ecommerce-microservice/user-service/internal/handler"
	framework "ecommerce-microservice/user-service/internal/infrastructure"
	"ecommerce-microservice/user-service/internal/repository"
	"ecommerce-microservice/user-service/internal/usecase"
)

func main() {

	// Framework / infrastructure
	db, err := framework.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repository
	userRepo := repository.NewSQLiteUserRepository(db)

	// Usecase
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Handler
	userHandler := handler.NewUserHandler(userUsecase)

	// Infrastructure/Framework
	router := framework.NewRouter(userHandler)

	// ========== Graceful Shutdown ===========
	// HTTP Server
	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		log.Println("User service running on :8081")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	log.Println("Shutdown signal received...")

	// Give active requests some time to finish
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shutdown HTTP server
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
