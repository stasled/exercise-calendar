package main

import (
	"context"
	"flag"
	"mycalendar/config"
	"mycalendar/internal/api/proto"
	"mycalendar/internal/storage/postgres"
	"mycalendar/internal/usecase"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	path := flag.String("config", "./config/", "Path to config folder")
	flag.Parse()

	logger.Debug("read config")
	config.GetConfig(*path)

	ctx := context.Background()

	dbClient := postgres.NewClient(ctx)
	storage := postgres.NewStorage(dbClient, logger)
	service := usecase.NewService(storage)

	logger.Info("run http server")
	proto.RunServer(ctx, service, logger)
}
