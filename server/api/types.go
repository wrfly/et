package api

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
