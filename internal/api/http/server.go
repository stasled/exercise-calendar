package http

import (
	"context"
	"log"
	"mycalendar/internal/api"
	"net/http"
	"os"
	"os/signal"
)

func RunServer(ctx context.Context, service api.Service, host, port string) {
	eh := NewEvents(ctx, service)

	sm := http.NewServeMux()
	sm.Handle("/", eh)

	s := http.Server{
		Addr:    host + ":" + port,
		Handler: sm,
	}

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
	log.Println("got signal:", sig)

	s.Shutdown(ctx)
}
