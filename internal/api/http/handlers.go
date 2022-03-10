package http

import (
	"context"
	"mycalendar/internal/api"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

// @title 		Calendar app
// @version 	1.0
// @description API Server for Calendar application

// @host 		localhost:8000
// @BasePath 	/

func NewHandler(ctx context.Context, service api.Service, logger *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://127.0.0.1:8000/swagger/doc.json"),
	))

	c := NewController(ctx, service, logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.MethodFunc("GET", "/events", c.GetAll)
		r.MethodFunc("POST", "/events", c.AddEvent)
		r.MethodFunc("PUT", "/events/{id}", c.UpdateEvent)
		r.MethodFunc("DELETE", "/events/{id}", c.DeleteEvent)
	})

	return r
}
