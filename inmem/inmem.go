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

type userNotificationRepository struct {
	mtx           sync.RWMutex
	notifications map[notification.ID]*notification.UserNotificationInfo
}

func (r *userNotificationRepository) Store(n *notification.UserNotificationInfo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.notifications[n.ID] = n
	return nil
}

// NewUserNotificationRepository returns a new instance of a in-memory notification repository.
func NewUserNotificationRepository() notification.UserNotificationRepository {
	return &userNotificationRepository{
		notifications: make(map[notification.ID]*notification.UserNotificationInfo),
	}
}

type subscriptionRepsoitory struct {
	mtx           sync.RWMutex
	subscriptions map[notification.ID]*notification.SubscriptionInfo
}

func (r *subscriptionRepsoitory) Insert(s *notification.SubscriptionInfo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.subscriptions[s.ID] = s
	return nil
}

func (r *subscriptionRepsoitory) GetAll(notificationName string) []notification.SubscriptionInfo {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	subscriptions := make([]notification.SubscriptionInfo, 0)
	for _, v := range r.subscriptions {
		if v.NotificationName == notificationName {
			subscriptions = append(subscriptions, *v)
		}
	}

	return subscriptions
}

// NewSubscriptionRepository returns a new instance of a in-memory notification repository.
func NewSubscriptionRepository() notification.SubscriptionRepository {
	return &subscriptionRepsoitory{
		subscriptions: make(map[notification.ID]*notification.SubscriptionInfo),
	}
}
