package api

import (
	"context"
	"mycalendar/internal/entity"
)

type Service interface {
	CreateEvent(ctx context.Context, title, startAt, endAt string) error
	UpdateEvent(ctx context.Context, e *entity.Event, title, startAt, endAt string) error
	DeleteEvent(ctx context.Context, id string) error
	GetEventByID(ctx context.Context, id string) (*entity.Event, error)
	GetEvents(ctx context.Context) (map[string]*entity.Event, error)
}
