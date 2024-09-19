package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	config "incexp-service/internal/pkg/load"
	pq "incexp-service/internal/pkg/postgres"
	service "incexp-service/internal/service"
	"incexp-service/logger"
	router "incexp-service/internal/pkg/register-service"
	producer "incexp-service/internal/kafka/producer"
)

func main() {

	logger.InitLog()

	cfg, err := config.Load("./config/config.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Configuration loaded")

	db, err := pq.ConnectDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Connected to Postgresql")

	writer, err := producer.NewProducerInit(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Producer started succesfully")
	defer writer.Close()

	repo := pq.NewIncExpRepository(db)
	srv := service.NewIncExpService(repo, *writer)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := router.NewGRPCIncExpService(srv)
		if err := r.RUN(cfg); err != nil {
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
