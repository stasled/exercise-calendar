package entity

type Event struct {
	Id      string `json:"id,omitempty"`
	Title   string `json:"title"`
	StartAt string `json:"start_at"`
	EndAt   string `json:"end_at"`
}
