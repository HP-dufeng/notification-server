package inmem

import (
	"sync"

	"github.com/fengdu/notification-server/core/notifications"
)

// NewNotificationInfoRepository returns a new instance of a in-memory notification repository.
func NewNotificationInfoRepository() notifications.NotificationInfoRepository {
	return &notificationRepository{
		data: make(map[notifications.ID]*notifications.NotificationInfo),
	}
}

// NewUserNotificationInfoRepository returns a new instance of a in-memory notification repository.
func NewUserNotificationInfoRepository() notifications.UserNotificationInfoRepository {
	return &userNotificationRepository{
		data: make(map[notifications.ID]*notifications.UserNotificationInfo),
	}
}

// NewSubscriptionInfoRepository returns a new instance of a in-memory notification repository.
func NewSubscriptionInfoRepository() notifications.SubscriptionInfoRepository {
	return &subscriptionRepsoitory{
		data: make(map[notifications.ID]*notifications.SubscriptionInfo),
	}
}

type notificationRepository struct {
	mtx  sync.RWMutex
	data map[notifications.ID]*notifications.NotificationInfo
}

func (r *notificationRepository) Insert(n *notifications.NotificationInfo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.data[n.ID] = n
	return nil
}

func (r *notificationRepository) Find(id notifications.ID) (*notifications.NotificationInfo, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if val, ok := r.data[id]; ok {
		return val, nil
	}

	return nil, notifications.ErrUnknown
}

type userNotificationRepository struct {
	mtx  sync.RWMutex
	data map[notifications.ID]*notifications.UserNotificationInfo
}

func (r *userNotificationRepository) Insert(n *notifications.UserNotificationInfo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.data[n.ID] = n
	return nil
}

func (r *userNotificationRepository) FindAll() []*notifications.UserNotificationInfo {
	s := make([]*notifications.UserNotificationInfo, 0, len(r.data))
	for _, v := range r.data {
		s = append(s, v)
	}

	return s
}

type subscriptionRepsoitory struct {
	mtx  sync.RWMutex
	data map[notifications.ID]*notifications.SubscriptionInfo
}

func (r *subscriptionRepsoitory) Insert(s *notifications.SubscriptionInfo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.data[s.ID] = s
	return nil
}

func (r *subscriptionRepsoitory) GetAll(notificationName string) []notifications.SubscriptionInfo {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	subscriptions := []notifications.SubscriptionInfo{}
	for _, v := range r.data {
		if v.NotificationName == notificationName {
			subscriptions = append(subscriptions, *v)
		}
	}

	return subscriptions
}
