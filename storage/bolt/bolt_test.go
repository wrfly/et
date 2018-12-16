package bolt

import (
	"testing"
	"time"

	"github.com/wrfly/et/types"
)

func TestBoltStorage(t *testing.T) {
	db, err := New("/tmp")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ID := "taskID" + time.Now().String()
	task := types.Task{
		ID:       ID,
		Comments: "test comments",
	}
	db.SaveTask(task)

	taskFound, err := db.FindTask(ID)
	if err != nil {
		t.Fatal(err)
	}
	if taskFound.Comments != task.Comments {
		t.Error("comments not equal")
	}

	for i := 0; i < 5; i++ {
		nID := "notificationID" + time.Now().String()
		n := types.Notification{
			ID:     nID,
			TaskID: ID,
			Event: types.OpenEvent{
				Time: time.Now(),
			},
		}
		if err := db.SaveNotification(n); err != nil {
			t.Error(err)
		}
	}

	ns, err := db.FindNotification(ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(ns) != 5 {
		t.Error("ns != 5")
	}
}
