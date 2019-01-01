package api

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/wrfly/et/types"
)

func printSum(bytes []byte) string {
	hasher := md5.New()
	hasher.Write(bytes)

	ID := ""
	sum := fmt.Sprintf("%x", hasher.Sum(nil))

	for i := 0; i < len(sum); i += 5 {
		end := i + 5
		if ID == "" {
			ID = fmt.Sprintf("%s", sum[i:end])
		} else {
			if end > len(sum) {
				end = len(sum)
			}
			ID = fmt.Sprintf("%s-%s", ID, sum[i:end])
		}
	}

	return ID
}

func genTaskID(t *types.Task) {
	t.ID = printSum([]byte(
		fmt.Sprint(t.NotifyTo, t.Comments, time.Now().UnixNano()),
	))
}

func genNotificationID(n types.Notification) string {
	return printSum([]byte(
		fmt.Sprint(n.TaskID, n.Event.IP, time.Now().UnixNano()),
	))
}

func newNotification(task types.Task, ip, ua string) types.Notification {
	n := types.Notification{
		TaskID: task.ID,
		Event: types.OpenEvent{
			IP:   ip,
			UA:   ua,
			Time: time.Now().Add(task.Offset),
		},
	}
	n.ID = genNotificationID(n)
	return n
}
