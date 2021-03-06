package bolt

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

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
		if err := b.Put(t.Key(), t.Value()); err != nil {
			return err
		}

		// update status
		statusB := getSBucket(tx)
		bs := statusB.Get(keyTaskSubmit)
		num, _ := strconv.ParseUint(string(bs), 10, 64)
		if err := statusB.Put(keyTaskSubmit,
			[]byte(fmt.Sprint(num+1))); err != nil {
			return err
		}

		bs = statusB.Get(keyToday(keyTaskSubmit))
		num, _ = strconv.ParseUint(string(bs), 10, 64)
		if err := statusB.Put(keyToday(keyTaskSubmit),
			[]byte(fmt.Sprint(num+1))); err != nil {
			return err
		}

		return nil
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

		// update status
		statusB := getSBucket(tx)
		bs := statusB.Get(keyNotification)
		num, _ := strconv.ParseUint(string(bs), 10, 64)
		err := statusB.Put(keyNotification, []byte(fmt.Sprint(num+1)))
		if err != nil {
			return err
		}

		bs = statusB.Get(keyToday(keyNotification))
		num, _ = strconv.ParseUint(string(bs), 10, 64)
		err = statusB.Put(keyToday(keyNotification), []byte(fmt.Sprint(num+1)))
		if err != nil {
			return err
		}

		relationBucket := getRBucket(tx)
		bs = relationBucket.Get([]byte(n.TaskID))
		notifications := make([]string, 0, 10)
		if len(bs) != 0 {
			if err := json.Unmarshal(bs, &notifications); err != nil {
				return err
			}
		}
		notifications = append(notifications, n.ID)
		bs, err = json.Marshal(notifications)
		if err != nil {
			return err
		}

		return relationBucket.Put([]byte(n.TaskID), bs)
	})
}

// FindNotification of the specific task ID
func (b *boltStorage) FindNotification(taskID string) ([]types.Notification, error) {
	ns := make([]types.Notification, 0, 10)
	ids := make([]string, 0, 10)

	if err := b.db.View(func(tx *bolt.Tx) error {
		relationBucket := getRBucket(tx)
		bs := relationBucket.Get([]byte(taskID))
		if len(bs) == 0 {
			return nil
		}
		if err := json.Unmarshal(bs, &ids); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return nil, storage.ErrNoNotification
	}

	if err := b.db.View(func(tx *bolt.Tx) error {
		notificationBucket := getNBucket(tx)
		for _, id := range ids {
			bs := notificationBucket.Get([]byte(id))
			n := types.Notification{}
			if err := json.Unmarshal(bs, &n); err != nil {
				return err
			}
			ns = append(ns, n)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Slice(ns, func(i, j int) bool { return ns[i].Event.Time.After(ns[j].Event.Time) })
	return ns, nil
}

func (b *boltStorage) Close() error {
	return b.db.Close()
}

func (b *boltStorage) Status() types.ServiceStatus {
	var (
		taskSubmit, notification           uint64
		taskSubmitDaily, notificationDaily uint64
	)
	b.db.View(func(tx *bolt.Tx) error {
		b := getSBucket(tx)
		bs := b.Get(keyTaskSubmit)
		taskSubmit, _ = strconv.ParseUint(string(bs), 10, 64)
		bs = b.Get(keyNotification)
		notification, _ = strconv.ParseUint(string(bs), 10, 64)

		bs = b.Get(keyToday(keyTaskSubmit))
		taskSubmitDaily, _ = strconv.ParseUint(string(bs), 10, 64)
		bs = b.Get(keyToday(keyNotification))
		notificationDaily, _ = strconv.ParseUint(string(bs), 10, 64)
		return nil
	})
	return types.ServiceStatus{
		Daily: types.RuntimeStatus{
			TaskSubmit:   taskSubmitDaily,
			Notification: notificationDaily,
		},
		Total: types.RuntimeStatus{
			TaskSubmit:   taskSubmit,
			Notification: notification,
		},
	}
}
