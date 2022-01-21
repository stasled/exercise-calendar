package usecase

import (
	"context"
	"errors"
	"mycalendar/internal/entity"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) *service {
	return &service{storage: storage}
}

var errorEmptyEventField = errors.New("one or more event fields are empty")
var errorEmptyEventId = errors.New("id field is empty")

func (s *service) CreateEvent(ctx context.Context, title, startAt, endAt string) error {
	if title == "" || startAt == "" || endAt == "" {
		return errorEmptyEventField
	}

	e := &entity.Event{
		Title:   title,
		StartAt: startAt,
		EndAt:   endAt,
	}
	return s.storage.Add(ctx, e)
}

func (s *service) UpdateEvent(ctx context.Context, e *entity.Event, title, startAt, endAt string) error {
	if title == "" && startAt == "" && endAt == "" {
		return errorEmptyEventField
	}

	if title != "" {
		e.Title = title
	}
	if startAt != "" {
		e.StartAt = startAt
	}
	if endAt != "" {
		e.EndAt = endAt
	}

	return s.storage.Update(ctx, e)
}

func (s *service) DeleteEvent(ctx context.Context, id string) error {
	if id == "" {
		return errorEmptyEventId
	}
	return s.storage.Delete(ctx, id)
}

func (s *service) GetEvents(ctx context.Context) (map[string]*entity.Event, error) {
	return s.storage.GetAll(ctx), nil
}

func (s *service) GetEventByID(ctx context.Context, id string) (*entity.Event, error) {
	if id == "" {
		return nil, errorEmptyEventId
	}
	return s.storage.GetByID(ctx, id)
}
