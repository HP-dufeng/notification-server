package notification

// UserNotificationInfo is the central class in the domain model.
type UserNotificationInfo struct {
	ID             ID
	UserID         int64
	NotificationID ID
	State          UserNotificationState
	CreationTime   int64
}

// UserNotificationRepository provides access a user_notification store.
type UserNotificationRepository interface {
}

// UserNotificationState describes state of notification.
type UserNotificationState int8
