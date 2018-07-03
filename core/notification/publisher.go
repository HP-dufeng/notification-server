package notification

type NotificationDto struct {
	NotificationName string
	Severity         Severity
	UserIds          []int64
	ExcludedUserIds  []int64
	Data             string
}

type UserNotificationDto struct {
	UserID                int64
	UserNotificationState UserNotificationState
	NotificationName      string
	Severity              Severity
	Data                  string
}

// Publisher used to distribute notifications to users.
type Publisher interface {
	Distribute(dto NotificationDto) error
}

type publisher struct {
	store            Store
	realtimeNotifier RealTimeNotifier
}

func (d *publisher) Distribute(dto NotificationDto) error {
	notification := &NotificationInfo{
		ID:               NextID(),
		NotificationName: dto.NotificationName,
		Severity:         dto.Severity,
		Data:             dto.Data,
	}
	err := d.store.InsertNotification(notification)
	if err != nil {
		return err
	}

	users := d.getUsers(&dto)

	userNotifications := d.SaveUserNotifications(users, notification)

	err = d.realtimeNotifier.SendNotifications(userNotifications)

	return err
}

func (d *publisher) getUsers(notification *NotificationDto) []int64 {
	users := make(map[int64]int64)
	for _, v := range notification.UserIds {
		users[v] = v
	}

	subscriptions := d.store.GetSubscriptions(notification.NotificationName)
	for _, v := range subscriptions {
		users[v.UserID] = v.UserID
	}

	for _, v := range notification.ExcludedUserIds {
		if u, ok := users[v]; ok {
			delete(users, u)
		}
	}

	validUsers := make([]int64, len(users))
	for _, v := range users {
		validUsers = append(validUsers, v)
	}
	return validUsers
}

func (d *publisher) SaveUserNotifications(userIds []int64, notification *NotificationInfo) []*UserNotificationDto {
	dtos := make([]*UserNotificationDto, len(userIds))
	for _, v := range userIds {
		d.store.InsertUserNotification(NewUserNotificationInfo(v, notification.ID))
		dtos = append(dtos, &UserNotificationDto{
			UserID:                v,
			UserNotificationState: Unread,
			NotificationName:      notification.NotificationName,
			Severity:              notification.Severity,
			Data:                  notification.Data,
		})
	}

	return dtos
}

// NewPublisher returns a new instance of a  Publisher interface.
func NewPublisher(store Store, realtimeNotifier RealTimeNotifier) Publisher {
	return &publisher{
		store,
		realtimeNotifier,
	}
}
