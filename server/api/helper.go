package api

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"strings"
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

func newBadge(w io.Writer, b *template.Template, name string, value interface{}) {
	b.Execute(w, map[string]interface{}{
		"name":      name,
		"len_name":  len(name)*50 + 1,
		"value":     fmt.Sprint(value),
		"len_value": len(fmt.Sprint(value))*70 + 1,
	})
}

func blueBadge(w io.Writer, name string, value interface{}) {
	newBadge(w, badgeBlue, name, value)
}

func greenBadge(w io.Writer, name string, value interface{}) {
	newBadge(w, badgeGreen, name, value)
}

func taskLinkToID(link string) string {
	x := strings.Split(link, "/")
	return strings.Split(x[len(x)-1], ".")[0]
}
