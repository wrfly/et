package bolt

import (
	"github.com/boltdb/bolt"
	"github.com/wrfly/et/storage"
	"github.com/wrfly/et/types"
)

type boltStorage struct {
	db *bolt.DB
}

// SaveTask to boltDB
func (b *boltStorage) SaveTask(t *types.Task) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		b := getTBucket(tx)
		return b.Put(t.Key(), t.Value())
	})
}

// FindTask in boltDB
func (b *boltStorage) FindTask(ID string) (*types.Task, error) {
	t := &types.Task{}

	return t, b.db.View(func(tx *bolt.Tx) error {
		b := getTBucket(tx)
		bs := b.Get([]byte(ID))
		if bs == nil {
			return storage.ErrTaskNotFound
		}
		t.Unmarshal(bs)
		return nil
	})
}

// SaveNotification not only save to db but also update the task
func (b *boltStorage) SaveNotification(n types.Notification) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		b := getNBucket(tx)
		if err := b.Put(n.Key(), n.Value()); err != nil {
			return err
		}

		// TODO: use relation bucket
		// r := getRBucket(tx)
		// r.Get([]byte(n.TaskID))
		return nil
	})
}

// FindNotification of the specific task ID
func (b *boltStorage) FindNotification(ID string) ([]types.Notification, error) {
	ns := make([]types.Notification, 0, 100)

	// TODO: need to use relation bucket for this
	return ns, b.db.View(func(tx *bolt.Tx) error {
		b := getNBucket(tx)
		b.ForEach(func(k, v []byte) error {
			n := types.Notification{}
			n.Unmarshal(v)
			if n.TaskID == ID {
				ns = append(ns, n)
			}
			return nil
		})
		if len(ns) == 0 {
			return storage.ErrNoNotification
		}
		return nil
	})
}

func (b *boltStorage) Close() error {
	return b.db.Close()
}
