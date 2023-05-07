package main

import (
	"bot/config"
	"bot/internal/bot"
	"bot/internal/storage"
	"bot/internal/usecase"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	store, err := storage.New()
	if err != nil {
		log.Fatalf("storage error: %s", err)
	}

	logic := usecase.New(store)

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("zap error: %s", err)
	}

	b, err := bot.New(cfg.Key, logic, logger)
	if err != nil {
		log.Fatalf("bot error: %s", err)
	}

	go func() {
		err := b.Start()
		if err != nil {
			log.Fatalf("bot error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutdown Server ...")

	b.Stop()
}
