package api

import (
	"time"

	"github.com/wrfly/et/types"
)

type taskRequest struct {
	NotifyTo string `json:"notifier"`
	Comments string `json:"comments"`
	Offset   int    `json:"offset"`
}

type taskResponse struct {
	Error     string `json:"err,omitempty"`
	TaskID    string `json:"id,omitempty"`
	TrackLink string `json:"link,omitempty"`
}

type taskGetResponse struct {
	State    string            `json:"state,omitempty"`
	Comments string            `json:"comments,omitempty"`
	Submit   time.Time         `json:"submitAt,omitempty"`
	Events   []types.OpenEvent `json:"events,omitempty"`
}
