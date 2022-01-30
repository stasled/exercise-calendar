package api

import (
	"context"
	"mycalendar/internal/entity"
	"time"
)

type Service interface {
	CreateEvent(ctx context.Context, title string, startAt, endAt time.Time) error
	UpdateEvent(ctx context.Context, e entity.Event, title string, startAt, endAt time.Time) error
	DeleteEvent(ctx context.Context, id int) error
	GetEventByID(ctx context.Context, id int) (entity.Event, error)
	GetEvents(ctx context.Context) (map[int]entity.Event, error)
}
