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
	Offset   time.Duration

	// task state
	State State
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

func (t *Task) BadState() bool {
	if t.State == StateNormal || t.State == StateResumed {
		return false
	}
	return true
}
