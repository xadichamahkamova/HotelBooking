package main

import (
	_ "api-gateway/docs"
	api "api-gateway/internal/http"
	producer "api-gateway/internal/kafka/producer"
	bookingService "api-gateway/internal/pkg/booking-service"
	hotelService "api-gateway/internal/pkg/hotel-service"
	config "api-gateway/internal/pkg/load"
	notifService "api-gateway/internal/pkg/notification-service"
	userService "api-gateway/internal/pkg/user-service"
	service "api-gateway/internal/service"
	"api-gateway/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	logger.InitLog()

	cfg, err := config.Load("config/config.yml")
	if err != nil {
		logger.Fatal("Failed to load configuration:", err)
	}
	logger.Info("Configuration loaded successfully")

	writer, err := producer.NewProducerInit(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Producer started succesfully")
	defer writer.Close()

	conn1, err := userService.DialWithUserService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to User Service:", err)
	}
	logger.Info("Connected to User Service")

	conn2, err := hotelService.DialWithHotelService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to Hotel Service:", err)
	}
	logger.Info("Connected to Hotel Service")

	conn3, err := bookingService.DialWithBookingService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to Booking Service:", err)
	}
	logger.Info("Connected to Booking Service")

	conn4, err := notifService.DialWithNotificationService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to Notification Service:", err)
	}
	logger.Info("Connected to Notification Service")

	clientService := service.NewServiceRepositoryClient(conn1, conn2, conn3, conn4)
	logger.Info("Service clients initialized")

	srv := api.NewGin(*clientService, *writer)
	addr := fmt.Sprintf(":%d", cfg.ApiGateway.Port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("Starting API Gateway on: ", addr)
		if err := srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
			logger.Fatal(err)
		}
	}()
	logger.Info("Starting API Gateway on address:", addr)

	signalReceived := <-sigChan
	logger.Info("Received signal:", signalReceived)

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("Server shutdown error: ", err)
	}
	logger.Info("Graceful shutdown complete.")
}
