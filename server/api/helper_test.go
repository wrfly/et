package api

import (
	"testing"

	"github.com/wrfly/et/types"
)

func TestGenTaskID(t *testing.T) {
	x := genTaskID(types.Task{
		NotifyTo: "null@kfd.me",
		Comments: "no comments",
	})
	t.Log(x)
}
