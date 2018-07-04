package inmem

import (
	"sync"

	"github.com/fengdu/notification-server/core/notifications"
)

type notificationRepository struct {
	mtx           sync.RWMutex
	notifications map[notifications.ID]*notifications.NotificationInfo
}

func (r *notificationRepository) Insert(n *notifications.NotificationInfo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.notifications[n.ID] = n
	return nil
}

func (r *notificationRepository) Find(id notifications.ID) (*notifications.NotificationInfo, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.notifications[id]; ok {
		return val, nil
	}

	return nil, notifications.ErrUnknown
}

// NewNotificationInfoRepository returns a new instance of a in-memory notification repository.
func NewNotificationInfoRepository() notifications.NotificationInfoRepository {
	return &notificationRepository{
		notifications: make(map[notifications.ID]*notifications.NotificationInfo),
	}
}

type userNotificationRepository struct {
	mtx           sync.RWMutex
	notifications map[notifications.ID]*notifications.UserNotificationInfo
}

func (r *userNotificationRepository) Insert(n *notifications.UserNotificationInfo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.notifications[n.ID] = n
	return nil
}

// NewUserNotificationInfoRepository returns a new instance of a in-memory notification repository.
func NewUserNotificationInfoRepository() notifications.UserNotificationInfoRepository {
	return &userNotificationRepository{
		notifications: make(map[notifications.ID]*notifications.UserNotificationInfo),
	}
}

type subscriptionRepsoitory struct {
	mtx           sync.RWMutex
	subscriptions map[notifications.ID]*notifications.SubscriptionInfo
}

func (r *subscriptionRepsoitory) Insert(s *notifications.SubscriptionInfo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.subscriptions[s.ID] = s
	return nil
}

func (r *subscriptionRepsoitory) GetAll(notificationName string) []notifications.SubscriptionInfo {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	subscriptions := make([]notifications.SubscriptionInfo, 0)
	for _, v := range r.subscriptions {
		if v.NotificationName == notificationName {
			subscriptions = append(subscriptions, *v)
		}
	}

	return subscriptions
}

// NewSubscriptionInfoRepository returns a new instance of a in-memory notification repository.
func NewSubscriptionInfoRepository() notifications.SubscriptionInfoRepository {
	return &subscriptionRepsoitory{
		subscriptions: make(map[notifications.ID]*notifications.SubscriptionInfo),
	}
}
