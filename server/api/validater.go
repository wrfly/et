package api

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/wrfly/et/limiter"
	"github.com/wrfly/et/types"
)

const (
	maxEmailLength   = 128
	maxCommentLength = 256
)

func validTask(t *types.Task) bool {
	if len(t.NotifyTo) > maxEmailLength {
		return false
	}
	x := strings.Split(t.NotifyTo, "@")
	if len(x) != 2 {
		return false
	}
	if len(x[1]) > maxEmailLength/2 {
		return false
	}

	if len(t.Comments) > maxCommentLength {
		return false
	}

	// TODO: only accept wellknown email address?

	return true
}

func checkTaskLimit(t *types.Task) bool {

	// check target email
	if err := limiter.Recv.Inc(t.NotifyTo); err != nil {
		logrus.Warnf("receiver [%s] limiter: %s", t.NotifyTo, err)
		return false
	}

	// check task send
	if err := limiter.Sent.Inc(t.ID); err != nil {
		limiter.Recv.Dec(t.NotifyTo)

		logrus.Warnf("task [%s] limiter: %s", t.ID, err)

		// change task state
		if t.State == types.StateNormal {
			t.State = types.StateStopped
		} else if t.State == types.StateResumed {
			t.State = types.StateTerminated
		}

		return false
	}

	return true
}
