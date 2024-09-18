package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	config "budget-service/internal/pkg/load"
	pq "budget-service/internal/pkg/postgres"
	service "budget-service/internal/service"
	"budget-service/logger"
	router "budget-service/internal/pkg/register-service"
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

	repo := pq.NewBudgetRepo(db)
	srv := service.NewBudgetService(repo)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := router.NewGRPCBudgetService(srv)
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
