package usecase

import (
	"context"
	"mycalendar/internal/entity"
	"mycalendar/internal/storage"
	"testing"
)

func TestService_AddEvent(t *testing.T) {
	type testCase struct {
		test        string
		id          string
		title       string
		startAt     string
		endAt       string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Add new event",
			id:          "1",
			title:       "New event",
			startAt:     "2021-11-25T15:04:00Z",
			endAt:       "2021-11-25T15:50:00Z",
			expectedErr: nil,
		},
		{
			test:        "Add new event with empty title",
			id:          "2",
			title:       "",
			startAt:     "2021-11-25T15:04:00Z",
			endAt:       "2021-11-25T15:50:00Z",
			expectedErr: errorEmptyEventField,
		},
		{
			test:        "Add new event with empty dates",
			id:          "2",
			title:       "New event",
			startAt:     "Z",
			endAt:       "",
			expectedErr: errorEmptyEventField,
		},
	}

	ctx := context.Background()
	storage := storage.NewEventLocalStorage()
	service := NewService(storage)
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := service.CreateEvent(ctx, tc.title, tc.startAt, tc.endAt)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}

func TestService_UpdateEvent(t *testing.T) {
	type testCase struct {
		test        string
		id          string
		title       string
		startAt     string
		endAt       string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Update event",
			id:          "1",
			title:       "Update event",
			startAt:     "2021-11-25T15:04:00Z",
			endAt:       "2021-11-25T15:50:00Z",
			expectedErr: nil,
		},
		{
			test:        "Update event without title",
			id:          "2",
			title:       "",
			startAt:     "2021-12-25T15:04:00Z",
			endAt:       "2021-12-25T15:50:00Z",
			expectedErr: nil,
		},
		{
			test:        "Update event by empty fields",
			id:          "2",
			title:       "",
			startAt:     "",
			endAt:       "",
			expectedErr: errorEmptyEventField,
		},
	}

	e := &entity.Event{
		Id:      "1",
		Title:   "Init event",
		StartAt: "2021-11-25T15:04:00Z",
		EndAt:   "2021-11-25T15:50:00Z",
	}

	ctx := context.Background()
	storage := storage.NewEventLocalStorage()
	service := NewService(storage)

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := service.UpdateEvent(ctx, e, tc.title, tc.startAt, tc.endAt)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}

// To Do test for Delete and GetEvents
