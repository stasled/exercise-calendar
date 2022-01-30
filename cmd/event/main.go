package main

import (
	"context"
	"flag"
	"go.uber.org/zap"
	"mycalendar/config"
	"mycalendar/internal/api/http"
	"mycalendar/internal/storage/postgres"
	"mycalendar/internal/usecase"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	path := flag.String("config", "./config/", "Path to config folder")
	flag.Parse()

	logger.Debug("read config")
	cfg := config.GetConfig(*path)

	ctx := context.Background()

	dbClient := postgres.NewClient(ctx, cfg)
	storage := postgres.NewStorage(dbClient, logger)
	service := usecase.NewService(storage)

	logger.Info("run http server")
	http.RunServer(ctx, service, cfg, logger)
}
