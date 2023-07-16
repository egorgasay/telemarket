package main

import (
	"bot/config"
	"bot/internal/bot"
	"bot/internal/handler"
	"bot/internal/storage"
	"bot/internal/usecase"
	api "github.com/egorgasay/telemarket-grpc/telemarket"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	store, err := storage.New(cfg.PathToItems)
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
		log.Println("Starting Bot ...")
		err := b.Start()
		if err != nil {
			log.Fatalf("bot error: %s", err)
		}
	}()

	h := handler.New(logic)

	grpcServer := grpc.NewServer()
	log.Println("Starting Telemarket ...")
	lis, err := net.Listen("tcp", "127.0.0.1:"+cfg.Port) // TODO: TO CONFIG
	if err != nil {
		log.Fatal("failed to listen", zap.Error(err))
	}
	api.RegisterTelemarketServer(grpcServer, h)

	// gRPC by default
	go func() {
		log.Println("Starting GRPC", lis.Addr())
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("grpcServer Serve: %v", err)
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutdown Bot ...")

	b.Stop()
}
