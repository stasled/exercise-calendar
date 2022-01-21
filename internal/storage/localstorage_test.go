package storage

import (
	"context"
	"mycalendar/internal/entity"
	"testing"
)

func TestAddEventToLocalStorage(t *testing.T) {
	type testCase struct {
		test        string
		title       string
		startAt     string
		endAt       string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Add new event",
			title:       "New event",
			startAt:     "2021-11-25T15:04:00Z",
			endAt:       "2021-11-25T15:50:00Z",
			expectedErr: nil,
		},
		{
			test:        "Add overlaping event",
			title:       "Overlaping event",
			startAt:     "2021-11-25T15:20:00Z",
			endAt:       "2021-11-25T15:30:00Z",
			expectedErr: ErrorEventBusy,
		},
		{
			test:        "Add empty event",
			expectedErr: ErrorEventData,
		},
	}

	ctx := context.Background()
	storage := NewEventLocalStorage()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			TestEvent := &entity.Event{
				Title:   tc.title,
				StartAt: tc.startAt,
				EndAt:   tc.endAt,
			}
			err := storage.Add(ctx, TestEvent)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}

// TO DO test for Update, Delete, GetAll, GetEventByID methods
