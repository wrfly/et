package types

import (
	"encoding/json"
	"time"
)

type OpenEvent struct {
	Time time.Time `json:"time,omitempty"`
	IP   string    `json:"ip,omitempty"`
	UA   string    `json:"ua,omitempty"`
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
