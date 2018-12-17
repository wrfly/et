package api

import (
	"strings"

	"github.com/wrfly/et/types"
)

const (
	maxEmailLength   = 128
	maxCommentLength = 256
)

func validTask(t types.Task) bool {
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
