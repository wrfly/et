package bolt

import (
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

// getRBucket returns notification bucket
func getRBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket([]byte(relationBucket))
}
