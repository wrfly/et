package types

import (
	// json lib is OK
	// performance is not that important to this project
	"encoding/json"
	"time"
)

type Task struct {
	ID string
	// comments, such as email subject, receiver, or some marks
	Comments string
	NotifyTo string
	// this field get from request, used to adjust notification time
	SubmitAt time.Time
	Adjust   time.Duration
	// opentime could change
	Opentimes int
}

func (t *Task) Key() []byte {
	return []byte(t.ID)
}

func (t *Task) Value() []byte {
	bs, _ := json.Marshal(t)
	return bs
}

func (t *Task) Unmarshal(bs []byte) {
	json.Unmarshal(bs, t)
}

type OpenEvent struct {
	Time time.Time
	IP   string
	UA   string
}

type Notification struct {
	ID     string
	TaskID string
	Event  OpenEvent
}

func (n *Notification) Key() []byte {
	return []byte(n.ID)
}

func (n *Notification) Value() []byte {
	bs, _ := json.Marshal(n)
	return bs
}

func (n *Notification) Unmarshal(bs []byte) {
	json.Unmarshal(bs, n)
}
