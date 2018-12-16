package api

import (
	"crypto/md5"
	"fmt"
	"time"
)

func genTaskID(mailTo string, comments string) string {
	hasher := md5.New()
	hasher.Write([]byte(
		fmt.Sprint(mailTo, comments, time.Now().UnixNano()),
	))

	taskID := ""
	sum := fmt.Sprintf("%x", hasher.Sum(nil))

	for i := 0; i < len(sum); i += 5 {
		end := i + 5
		if taskID == "" {
			taskID = fmt.Sprintf("%s", sum[i:end])
		} else {
			if end > len(sum) {
				end = len(sum)
			}
			taskID = fmt.Sprintf("%s-%s", taskID, sum[i:end])
		}
	}

	return taskID
}

func genNotificationID(salt string) string {
	return genTaskID("", salt)
}
