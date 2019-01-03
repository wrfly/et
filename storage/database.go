package storage

import (
	"errors"

	"github.com/wrfly/et/types"
)

// errors
var (
	ErrTaskNotFound   = errors.New("task not found")
	ErrNoNotification = errors.New("no notification found")
)

type Database interface {
	SaveTask(*types.Task) error
	FindTask(ID string) (*types.Task, error)
	// SaveNotification not only save to db but also update the task
	SaveNotification(types.Notification) error
	// find all notifications of the specific task ID
	FindNotification(taskID string) ([]types.Notification, error)
	Status() types.ServiceStatus
	Close() error
}
