package api

import "testing"

func TestGenTaskID(t *testing.T) {
	x := genTaskID("null@kfd.me", "no comments")
	t.Log(x)
}
