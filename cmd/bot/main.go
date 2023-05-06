package main

import (
	"bot/config"
	"bot/internal/bot"
	"bot/internal/storage"
	"bot/internal/usecase"
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

	store := storage.New()
	logic := usecase.New(store)

	b, err := bot.New(cfg.Key, logic)
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
