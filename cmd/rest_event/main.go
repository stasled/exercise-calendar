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

// @title Calendar API
// @version 1.0
// @description App Calendar server REST API.
// @BasePath /api/v1
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
	http.RunServer(ctx, service, logger)
}
