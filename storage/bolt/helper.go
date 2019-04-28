package bolt

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var (
	_taskBucket         = []byte("tasks")
	_notificationBucket = []byte("notifications")
	_relationBucket     = []byte("task->notification")
	_statusBucket       = []byte("status")
)

// getTBucket returns task bucket
func getTBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket(_taskBucket)
}

// getNBucket returns notification bucket
func getNBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket(_notificationBucket)
}

// getRBucket returns releationship bucket
func getRBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket(_relationBucket)
}

// getSBucket returns status bucket
func getSBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket(_statusBucket)
}

var (
	keyNotification = []byte("notification")
	keyTaskSubmit   = []byte("taskSubmit")
)

func keyToday(key []byte) []byte {
	return []byte(fmt.Sprintf("%s-%d",
		key, time.Now().Truncate(time.Hour*24).Unix()))
}
