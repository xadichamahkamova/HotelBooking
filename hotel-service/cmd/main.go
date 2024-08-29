package main

import (
	"context"
	config "hotel-service/internal/pkg/load"
	pq "hotel-service/internal/pkg/postgres"
	router "hotel-service/internal/pkg/register-service"
	hotelRepo "hotel-service/internal/repository"
	hotelService "hotel-service/internal/service"
	"hotel-service/logger"
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

	in := hotelRepo.NewHotelRepository(db)
	service := hotelService.NewHotelService(in)

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
