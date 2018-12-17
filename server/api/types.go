package api

import "time"

type taskRequest struct {
	NotifyTo  string    `json:"notifier"`
	Comments  string    `json:"comments"`
	LocalTime time.Time `json:"time"`
}

type taskResponse struct {
	Error     string `json:"err,omitempty"`
	TaskID    string `json:"id,omitempty"`
	TrackLink string `json:"link,omitempty"`
}
