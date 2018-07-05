package notifications

// Publisher used to distribute notifications to users.
type Publisher interface {
	Publish(notificationName string, message string, severity Severity, userIds []int64, excludedUserIds []int64) error
}

type publisher struct {
	store            Store
	realtimeNotifier RealTimeNotifier
}

// NewPublisher returns a new instance of a  Publisher interface.
func NewPublisher(store Store, realtimeNotifier RealTimeNotifier) Publisher {
	return &publisher{
		store,
		realtimeNotifier,
	}
}

func (d *publisher) Publish(notificationName string, message string, severity Severity, userIds []int64, excludedUserIds []int64) error {
	notification := &NotificationInfo{
		ID:               NextID(),
		NotificationName: notificationName,
		Severity:         severity,
		Data:             message,
	}
	err := d.store.InsertNotification(notification)
	if err != nil {
		return err
	}

	users := d.getUsers(notificationName, userIds, excludedUserIds)

	userNotifications := d.saveUserNotifications(users, notification)

	err = d.realtimeNotifier.SendNotifications(userNotifications)

	return err
}

func (d *publisher) getUsers(notificationName string, userIds []int64, excludedUserIds []int64) []int64 {
	users := make(map[int64]int64)
	for _, v := range userIds {
		users[v] = v
	}

	subscriptions := d.store.GetSubscriptions(notificationName)
	for _, v := range subscriptions {
		users[v.UserID] = v.UserID
	}

	for _, v := range excludedUserIds {
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

func (d *publisher) saveUserNotifications(userIds []int64, notification *NotificationInfo) []*UserNotificationDto {
	dtos := make([]*UserNotificationDto, 0, len(userIds))
	for _, v := range userIds {
		d.store.InsertUserNotification(NewUserNotificationInfo(v, notification.ID))
		dtos = append(dtos, &UserNotificationDto{
			UserID:                v,
			UserNotificationState: Unread.String(),
			NotificationName:      notification.NotificationName,
			Severity:              notification.Severity.String(),
			Data:                  notification.Data,
		})
	}

	return dtos
}
