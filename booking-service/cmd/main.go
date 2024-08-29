package main

import (
	"context"
	config "booking-service/internal/pkg/load"
	pq "booking-service/internal/pkg/postgres"
	router "booking-service/internal/pkg/register-service"
	bookingRepo "booking-service/internal/repository"
	bookingService "booking-service/internal/service"
	"booking-service/logger"
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
		logger.Fatal(err)
	}
	logger.Info("Configuration loaded")
	
	db, err := pq.ConnectDB(*cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Successfully connected to Postgres")

	in := bookingRepo.NewBookingRepository(db)
	service :=bookingService.NewBookingService(in)

	var wg sync.WaitGroup
 	wg.Add(1)
	go func() {
		defer wg.Done()
		r := router.NewGRPCHotelService(service)
		if err := r.RUN(*cfg); err != nil {
			logger.Fatal(err)
		}
	}()

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signChan

	logger.Info("Received signal: ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	router.GracefufShutdown()

	<-ctx.Done()
	logger.Info("Graceful shutdown complete.")
}
