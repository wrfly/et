package bolt

import (
	"fmt"
	"os"
	"path"

	"github.com/boltdb/bolt"
	"github.com/wrfly/et/storage"
)

// New storage
func New(dbRootPath string) (storage.Database, error) {
	f, err := os.Stat(dbRootPath)
	if err != nil {
		return nil, fmt.Errorf("get %s stat error: %s", dbRootPath, err)
	}
	if !f.IsDir() {
		return nil, fmt.Errorf("can not create boltDB, dbRootPath is not a dir")
	}

	dbFilePath := path.Join(dbRootPath, "email-tracker.db")
	db, err := bolt.Open(dbFilePath, 0600, nil)
	if err != nil {
		return nil, err
	}

	// create buckets
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(_taskBucket)
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists(_relationBucket)
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists(_statusBucket)
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists(_notificationBucket)
		return err
	}); err != nil {
		return nil, err
	}

	return &boltStorage{db}, nil
}
