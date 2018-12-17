package notify

import (
	"fmt"

	"github.com/wrfly/et/types"
)

type Notifier interface {
	Send(to, content string) (string, int, error)
}

const (
	subject = "[Email Tracker] knock knock, someone opens your email!"

	testReceiver = "null@kfd.me"

	NotifyTemplate = `Hey,
  Your email was opened!

  Comments: %s
  ---
    IP: %s
    UA: %s
    Date: %s
  ---

Best Regards,
https://track.kfd.me
`
)

var (
	constSender      = "Track Service"
	constSenderEmail = "track@kfd.me"
)

func ResetSender(name, email string) {
	constSender, constSenderEmail = name, email
}

// TestNotifier send email to null@kfd.me to check the API works
func TestNotifier(n Notifier) (string, error) {
	body, code, err := n.Send(testReceiver, "x")
	if err != nil {
		return body, err
	}
	if code != 200 {
		return body, nil
	}

	return body, nil
}

func NewContent(n types.Notification, comments string) string {
	return fmt.Sprintf(NotifyTemplate,
		comments,
		n.Event.IP,
		n.Event.UA,
		n.Event.Time.Format("2006-01-02 15:04:05"),
	)
}
