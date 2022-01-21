package http

import (
	"context"
	"mycalendar/internal/api"
	"net/http"
)

type Events struct {
	ctx     context.Context
	service api.Service
}

type Event struct {
	Id      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	StartAt string `json:"start_at,omitempty"`
	EndAt   string `json:"end_at,omitempty"`
}

func NewEvents(ctx context.Context, s api.Service) *Events {
	return &Events{
		ctx:     ctx,
		service: s,
	}
}

func (e *Events) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		e.GetAll(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		e.AddEvent(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		e.UpdateEvent(rw, r)
		return
	}

	if r.Method == http.MethodDelete {
		e.DeleteEvent(rw, r)
		return
	}
}
