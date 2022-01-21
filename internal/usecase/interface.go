package usecase

import (
	"context"
	"mycalendar/internal/entity"
)

type Storage interface {
	Add(ctx context.Context, event *entity.Event) error
	Update(ctx context.Context, event *entity.Event) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*entity.Event, error)
	GetAll(ctx context.Context) map[string]*entity.Event
}
