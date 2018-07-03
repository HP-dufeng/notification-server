package notification

// RealTimeNotifier to send real time notifications to users.
type RealTimeNotifier interface {
	SendNotifications(notifications []*UserNotificationDto) error
}
