package main

import (
	"context"
	notifRepo "notification-service/internal/email"
	config "notification-service/internal/pkg/load"
	router "notification-service/internal/pkg/register-service"
	notifService "notification-service/internal/service"
	"notification-service/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	logger.InitLog()

	cfg, err := config.Load("config/config.yml")
	if err != nil {
		logger.Error("Failed to load configuration: ", err)
	}
	logger.Info("Configuration loaded successfully")

	notif := notifRepo.NewNotificationRepo(*cfg)
	service := notifService.NewNotificationService(*notif)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := router.NewGRPCNotificationService(service)
		logger.Info("Starting gRPC Notification Service...")
		if err := r.RUN(*cfg); err != nil {
			logger.Fatal("Failed to start gRPC Notification Service: ", err)
		}
	}()

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signChan

	logger.Info("Received signal: ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Info("Initiating graceful shutdown...")
	router.GracefufShutdown()

	<-ctx.Done()
	logger.Info("Graceful shutdown complete.")
}
