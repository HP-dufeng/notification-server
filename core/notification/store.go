package notification

// Store used to store (persist) notifications.
type Store interface {
	InsertNotification(n *NotificationInfo) error
	InsertUserNotification(userNotificationInfo *UserNotificationInfo) error
	GetNotification(notificationID ID) (*NotificationInfo, error)
	GetSubscriptions(notificationName string) []SubscriptionInfo
}

type store struct {
	notificationRepository     Repository
	userNotificationRepository UserNotificationRepository
	subscriptionRespository    SubscriptionRepository
}

func (s *store) InsertNotification(n *NotificationInfo) error {
	return s.notificationRepository.Insert(n)
}

func (s *store) InsertUserNotification(u *UserNotificationInfo) error {
	err := s.userNotificationRepository.Insert(u)
	return err
}

func (s *store) GetNotification(notificationID ID) (*NotificationInfo, error) {
	return s.notificationRepository.Find(notificationID)

}

func (s *store) GetSubscriptions(notificationName string) []SubscriptionInfo {
	return s.subscriptionRespository.GetAll(notificationName)
}

// NewNotificationStore returns a new instance of a notification store.
func NewNotificationStore(notificationRepository Repository, userNotificationRepository UserNotificationRepository, subscriptionRepository SubscriptionRepository) Store {
	return &store{
		notificationRepository,
		userNotificationRepository,
		subscriptionRepository,
	}
}
