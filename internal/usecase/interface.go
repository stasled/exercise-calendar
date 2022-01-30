package usecase

import (
	"context"
	"mycalendar/internal/entity"
)

type Storage interface {
	Add(ctx context.Context, event entity.Event) error
	Update(ctx context.Context, event entity.Event) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (entity.Event, error)
	GetAll(ctx context.Context) (map[int]entity.Event, error)
}
