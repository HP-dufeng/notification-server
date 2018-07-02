package inmem

import (
	"sync"

	"github.com/fengdu/notification-server/core/notification"
)

type notificationRepository struct {
	mtx           sync.RWMutex
	notifications map[notification.ID]*notification.Notification
}

func (r *notificationRepository) Store(n *notification.Notification) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.notifications[n.ID] = n
	return nil
}

func (r *notificationRepository) Find(id notification.ID) (*notification.Notification, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.notifications[id]; ok {
		return val, nil
	}

	return nil, notification.ErrUnknown
}

// NewNotificationRepository returns a new instance of a in-memory notification repository.
func NewNotificationRepository() notification.Repository {
	return &notificationRepository{
		notifications: make(map[notification.ID]*notification.Notification),
	}
}
