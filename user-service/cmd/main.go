package main

import (
	"context"
	config "user-service/internal/pkg/load"
	pq "user-service/internal/pkg/postgres"
	router "user-service/internal/pkg/register-service"
	userRepo "user-service/internal/repository"
	userService "user-service/internal/service"
	"user-service/logger"
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

	in := userRepo.NewUserRepository(db)
	service := userService.NewUserService(in)

	var wg sync.WaitGroup
 	wg.Add(1)
	go func() {
		defer wg.Done()
		r := router.NewGRPCUserService(service)
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
