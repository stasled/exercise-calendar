package http

import (
	"context"
	"fmt"
	"log"
	_ "mycalendar/docs/rest"
	"mycalendar/internal/api"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func RunServer(ctx context.Context, service api.Service, logger *zap.Logger) {
	logger.Debug("Register handler")
	eh := NewHandler(ctx, service, logger)

	s := http.Server{
		Addr:         fmt.Sprintf("%s:%d", viper.GetString("REST_SERVER_HOST"), viper.GetInt("REST_SERVER_PORT")),
		Handler:      eh,
		ReadTimeout:  viper.GetDuration("REST_SERVER_TIME_READ") * time.Second,
		WriteTimeout: viper.GetDuration("REST_SERVER_TIME_WRITE") * time.Second,
		IdleTimeout:  viper.GetDuration("REST_SERVER_TIME_IDLE") * time.Second,
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
