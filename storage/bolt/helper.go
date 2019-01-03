package bolt

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

// getTBucket returns task bucket
func getTBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket([]byte(taskBucket))
}

// getNBucket returns notification bucket
func getNBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket([]byte(notificationBucket))
}

// getRBucket returns releationship bucket
func getRBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket([]byte(relationBucket))
}

// getSBucket returns status bucket
func getSBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket([]byte(statusBucket))
}

var (
	keyNotification = []byte("notification")
	keyTaskSubmit   = []byte("taskSubmit")
)

func keyToday(key []byte) []byte {
	return []byte(fmt.Sprintf("%s-%d",
		key, time.Now().Truncate(time.Hour*24).Unix()))
}
