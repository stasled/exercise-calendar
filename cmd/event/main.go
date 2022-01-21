package main

import (
	"context"
	"flag"
	"mycalendar/config"
	"mycalendar/internal/api/http"
	"mycalendar/internal/api/proto"
	"mycalendar/internal/storage"
	"mycalendar/internal/usecase"
	"time"
)

func main() {
	path := flag.String("config", "./config/dev.json", "Path to config file")
	flag.Parse()

	cfg := config.GetConfig(*path)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	ls := storage.NewEventLocalStorage()
	service := usecase.NewService(ls)

	http.RunServer(ctx, service, cfg.HttpServer.Host, cfg.HttpServer.Port)
	proto.RunServer(ctx, service, cfg.GrpcServer.Host, cfg.GrpcServer.Port)
}
