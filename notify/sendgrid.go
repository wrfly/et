package notify

// send emails with sendgrid

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sgNotifier struct {
	cli       *sendgrid.Client
	fromName  string
	fromEmail string
}

func (n *sgNotifier) Send(toMail, content string) (string, int, error) {
	from := &mail.Email{
		Name:    constSender,
		Address: constSenderEmail,
	}
	to := &mail.Email{
		Name:    toMail,
		Address: toMail,
	}

	m := mail.NewV3MailInit(
		from,
		subject,
		to,
		mail.NewContent("text/plain", content),
	)

	m.SetTrackingSettings(&mail.TrackingSettings{
		ClickTracking: mail.NewClickTrackingSetting().
			SetEnable(false).
			SetEnableText(false),
		OpenTracking: mail.NewOpenTrackingSetting().
			SetEnable(false),
		Footer: mail.NewFooterSetting().
			SetEnable(false),
	})

	resp, err := n.cli.Send(m)
	if err != nil {
		return "", -1, err
	}
	return resp.Body, resp.StatusCode, nil
}

func NewSendgridNotifier(API string) Notifier {
	cli := sendgrid.NewSendClient(API)
	return &sgNotifier{
		cli: cli,
	}
}
