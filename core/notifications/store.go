package notifications

// Store used to store (persist) notifications.
type Store interface {
	InsertNotification(n *NotificationInfo) error
	InsertUserNotification(userNotificationInfo *UserNotificationInfo) error
	GetNotification(notificationID ID) (*NotificationInfo, error)
	GetUserNotifications(userID int64) []*UserNotificationInfo
	GetSubscriptions(notificationName string) []SubscriptionInfo
}

type Repositories struct {
	NotificationInfoRepository
	UserNotificationInfoRepository
	SubscriptionInfoRepository
}

// NewNotificationStore returns a new instance of a notification store.
func NewNotificationStore(repositories Repositories) Store {
	return &store{
		repositories.NotificationInfoRepository,
		repositories.UserNotificationInfoRepository,
		repositories.SubscriptionInfoRepository,
	}
}

type store struct {
	notificationInfoRepository     NotificationInfoRepository
	userNotificationInfoRepository UserNotificationInfoRepository
	subscriptionInfoRespository    SubscriptionInfoRepository
}

func (s *store) InsertNotification(n *NotificationInfo) error {
	return s.notificationInfoRepository.Insert(n)
}

func (s *store) InsertUserNotification(u *UserNotificationInfo) error {
	err := s.userNotificationInfoRepository.Insert(u)
	return err
}

func (s *store) GetNotification(notificationID ID) (*NotificationInfo, error) {
	return s.notificationInfoRepository.Find(notificationID)

}

func (s *store) GetUserNotifications(userID int64) []*UserNotificationInfo {
	userNotifications := make([]*UserNotificationInfo, 0)
	for _, v := range s.userNotificationInfoRepository.FindAll() {
		if v.UserID == userID {
			userNotifications = append(userNotifications, v)
		}
	}

	return userNotifications
}

func (s *store) GetSubscriptions(notificationName string) []SubscriptionInfo {
	return s.subscriptionInfoRespository.GetAll(notificationName)
}
