package inmem

import (
	"sync"

	"github.com/fengdu/notification-server/notification"
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

// NewNotificationRepository returns a new instance of a in-memory notification repository.
func NewNotificationRepository() notification.Repository {
	return &notificationRepository{
		notifications: make(map[notification.ID]*notification.Notification),
	}
}
