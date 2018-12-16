package types

import (
	"time"
)

type Task struct {
	ID string
	// comments, such as email subject, receiver, or some marks
	Comments  string
	NotifyTo  string
	Opentimes int
}

type OpenEvent struct {
	Date time.Time
	IP   string
	UA   string
}

type Notification struct {
	ID    string
	Task  Task
	Event OpenEvent
}
