package entity

import "time"

type Event struct {
	Id      int       `json:"id,omitempty"`
	Title   string    `json:"title"`
	StartAt time.Time `json:"start_at"`
	EndAt   time.Time `json:"end_at"`
}
