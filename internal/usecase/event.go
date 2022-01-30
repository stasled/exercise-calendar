package usecase

import (
	"context"
	"errors"
	"mycalendar/internal/entity"
	"time"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) *service {
	return &service{storage: storage}
}

var errorEmptyEventField = errors.New("one or more event fields are empty")
var errorEmptyEventId = errors.New("id field is empty")

func (s *service) CreateEvent(ctx context.Context, title string, startAt, endAt time.Time) error {
	if title == "" {
		return errorEmptyEventField
	}

	e := entity.Event{
		Title:   title,
		StartAt: startAt,
		EndAt:   endAt,
	}
	return s.storage.Add(ctx, e)
}

func (s *service) UpdateEvent(ctx context.Context, e entity.Event, title string, startAt, endAt time.Time) error {
	if title == "" {
		return errorEmptyEventField
	}

	if title != "" {
		e.Title = title
	}
	e.StartAt = startAt
	e.EndAt = endAt

	return s.storage.Update(ctx, e)
}

func (s *service) DeleteEvent(ctx context.Context, id int) error {
	if id == 0 {
		return errorEmptyEventId
	}
	return s.storage.Delete(ctx, id)
}

func (s *service) GetEvents(ctx context.Context) (map[int]entity.Event, error) {
	return s.storage.GetAll(ctx)
}

func (s *service) GetEventByID(ctx context.Context, id int) (entity.Event, error) {
	if id == 0 {
		return entity.Event{}, errorEmptyEventId
	}
	return s.storage.GetByID(ctx, id)
}
