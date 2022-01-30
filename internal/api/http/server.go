package http

import (
	"context"
	"fmt"
	"log"
	"mycalendar/config"
	_ "mycalendar/docs"
	"mycalendar/internal/api"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
)

func RunServer(ctx context.Context, service api.Service, cfg *config.Config, logger *zap.Logger) {
	logger.Debug("Register handler")
	eh := NewHandler(ctx, service, logger)

	s := http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
		Handler:      eh,
		ReadTimeout:  cfg.HttpServer.TimeoutRead * time.Second,
		WriteTimeout: cfg.HttpServer.TimeoutWrite * time.Second,
		IdleTimeout:  cfg.HttpServer.TimeoutIdle * time.Second,
	}

	logger.Info("Run http server")
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatalf("unable to run http server: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	logger.Info(fmt.Sprintf("got signal:", sig))

	s.Shutdown(ctx)
}
