package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"notification-service/internal/email"
	consumer "notification-service/internal/kafka/consumer"
	"notification-service/logger"

	handler "notification-service/internal/http/handler"
	config "notification-service/internal/pkg/load"
)

func main() {

	logger.InitLog()

	cfg, err := config.Load("./config/config.yaml")
	if err != nil {
		logger.Fatal("Failed to load configuration: ", err)
	}
	logger.Info("Configuration loaded successfully")

	reader, err := consumer.NewConsumeInit(cfg)
	if err != nil {
		logger.Fatal("Failed to initialize Kafka consumer: ", err)
	}
	logger.Info("Kafka client created successfully")

	sender:= email.NewSMTPRouter(*cfg)
	defer func() {
		if err := reader.Close(); err != nil {
			logger.Error("Failed to close Kafka consumer: ", err)
		} else {
			logger.Info("Kafka consumer closed successfully")
		}
	}()

	http.HandleFunc("/ws", handler.HandleWebSocket)

	var wg sync.WaitGroup
	wg.Add(2)
	h := handler.HandlerST{
		Sender: *sender,
	}

	go func() {
		defer wg.Done()
		h.ConsumeMessage(*reader)
	}()

	addr := fmt.Sprintf(":%s", cfg.Service.Port)

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:      addr,
		Handler:   http.DefaultServeMux,
		TLSConfig: tlsConfig,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer wg.Done()
		logger.Info("Notification service listening on ", addr)
		if err := srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
			logger.Fatal(err)
		}
	}()
	wg.Wait()

	signalReceived := <-sigChan
	logger.Info("Received signal:", signalReceived)

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("Server shutdown error: ", err)
	}
	logger.Info("Graceful shutdown complete.")
}

