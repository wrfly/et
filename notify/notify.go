package notify

type Notifier interface {
	Send(to, content string) (string, int, error)
}

const (
	subject = "[Email Tracker] knock knock, someone opens your email!"

	testReceiver = "null@kfd.me"
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
