package storage

import (
	"context"
	"errors"
	"mycalendar/internal/entity"
	"sync"
	"time"

	"github.com/google/uuid"
)

type EventLocal struct {
	id      string
	title   string
	startAt time.Time
	endAt   time.Time
}

type EventLocalStorage struct {
	events map[string]*EventLocal
	mu     *sync.RWMutex
}

var ErrorEventBusy = errors.New("this time is busy with another event")
var ErrorEventData = errors.New("the data is incorrect")
var ErrNotFound = errors.New("not found")

func NewEventLocalStorage() *EventLocalStorage {
	return &EventLocalStorage{
		events: make(map[string]*EventLocal),
		mu:     new(sync.RWMutex),
	}
}

func (e *EventLocalStorage) Add(ctx context.Context, event *entity.Event) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	localEvent, err := toEventLocal(event)
	if err != nil {
		return ErrorEventData
	}

	timeIsBusy := checkTimeIsBusy(e, localEvent)
	if timeIsBusy {
		return ErrorEventBusy
	}

	id := uuid.New().String()
	localEvent.id = id
	e.events[id] = localEvent

	return nil
}

func (e *EventLocalStorage) Update(ctx context.Context, event *entity.Event) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	el, err := toEventLocal(event)
	if err != nil {
		return ErrorEventData
	}
	e.events[event.Id] = el

	return nil
}

func (e *EventLocalStorage) Delete(ctx context.Context, id string) error {
	e.mu.Lock()
	delete(e.events, id)
	e.mu.Unlock()

	return nil
}

func (e *EventLocalStorage) GetAll(ctx context.Context) map[string]*entity.Event {
	e.mu.RLock()
	events := make(map[string]*entity.Event)
	for index, event := range e.events {
		events[index] = toEvent(event, time.RFC3339)
	}
	e.mu.RUnlock()

	return events
}

func (e *EventLocalStorage) GetByID(ctx context.Context, id string) (*entity.Event, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	for index, event := range e.events {
		if index == id {
			e := toEvent(event, time.RFC3339)
			return e, nil
		}
	}
	return nil, ErrNotFound
}

func toEventLocal(e *entity.Event) (*EventLocal, error) {
	startAt, err := toDateTimeFormat(e.StartAt, time.RFC3339)
	if err != nil {
		return nil, err
	}

	endAt, err := toDateTimeFormat(e.EndAt, time.RFC3339)
	if err != nil {
		return nil, err
	}

	return &EventLocal{
		id:      e.Id,
		title:   e.Title,
		startAt: startAt,
		endAt:   endAt,
	}, nil
}

func toEvent(e *EventLocal, format string) *entity.Event {
	return &entity.Event{
		Id:      e.id,
		Title:   e.title,
		StartAt: e.startAt.Format(format),
		EndAt:   e.endAt.Format(format),
	}
}

func toDateTimeFormat(dateTime string, format string) (time.Time, error) {
	dt, err := time.Parse(format, dateTime)
	if err != nil {
		return time.Time{}, err
	}

	return dt, nil
}

func checkTimeIsBusy(s *EventLocalStorage, event *EventLocal) bool {
	for _, s := range s.events {
		if (event.startAt.After(s.startAt) && event.startAt.Before(s.endAt)) ||
			(event.endAt.After(s.startAt) && event.endAt.Before(s.endAt)) ||
			event.startAt.Equal(s.startAt) || event.endAt.Equal(s.startAt) ||
			event.startAt.Equal(s.endAt) || event.endAt.Equal(s.endAt) {
			return true
		}
	}

	return false
}
