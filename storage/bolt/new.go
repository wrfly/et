package bolt

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/wrfly/et/storage"
)

const (
	taskBucket         = "tasks"
	notificationBucket = "notifications"
	relationBucket     = "task->notification"
	statusBucket       = "status"
)

func New(dbRoot string) (storage.Database, error) {
	if f, err := os.Stat(dbRoot); err != nil {
		return nil, err
	} else {
		if !f.IsDir() {
			return nil, fmt.Errorf("can not create boltDB, dbRoot is not a dir")
		}
	}

	db, err := bolt.Open(dbRoot+"/email-tracker.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	// create buckets
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(taskBucket))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(relationBucket))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(statusBucket))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(notificationBucket))
		return err
	}); err != nil {
		return nil, err
	}

	return &boltStorage{db}, nil
}
