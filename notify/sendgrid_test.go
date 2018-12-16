package notify

import (
	"os"
	"testing"
)

func TestSendGrid(t *testing.T) {
	api := os.Getenv("SENDGRID_API_KEY")
	if api == "" {
		t.Log("No API Key, skip test")
		return
	}
	n := NewSendgridNotifier(api)
	body, err := TestNotifier(n)
	if err != nil {
		t.Error(err)
	}
	if body != "" {
		t.Error(body)
	}
}

func TestNotify(t *testing.T) {
	api := os.Getenv("SENDGRID_API_KEY")
	if api == "" {
		t.Log("No API Key, skip test")
		return
	}
	n := NewSendgridNotifier(api)
	body, code, err := n.Send("wrtset@mail.com", "hello email tracker")
	if err != nil {
		t.Error(err)
	}
	if body != "" {
		t.Error(body)
	}
	if code != 0 {
		t.Logf("got code: %d", code)
	}
}
