package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ecommerce-microservice/api-gateway/proxy"
)

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	router := gin.Default()

	// Public
	router.POST("/api/login", login)

	// Protected
	protected := router.Group("/api/v1")
	//protected.Use(middleware.JWTAuth())  // Authorization with JWT

	protected.Any("/users/*path", proxy.UserServiceProxy())   // For users
	protected.Any("/orders/*path", proxy.OrderServiceProxy()) // For orders

	// ==========================
	// HTTP Server
	// ==========================

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {

		log.Println("API Gateway running on :8080")

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

	log.Println("API Gateway stopped")

}

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login endpoint",
	})
}
