package notifications

// UserNotificationDto represents a notification sent to a user.
type UserNotificationDto struct {
	UserID                int64
	UserNotificationState UserNotificationState
	NotificationName      string
	Severity              Severity
	Data                  string
}

// RealTimeNotifier to send real time notifications to users.
type RealTimeNotifier interface {
	SendNotifications(notifications []*UserNotificationDto) error
}

type nullRealTimeNotifier struct{}

// NewNullRealTimeNotifier null pattern implementation of RealTimeNotifier interface.
func NewNullRealTimeNotifier() RealTimeNotifier {
	return nullRealTimeNotifier{}
}

func (nullRealTimeNotifier) SendNotifications(notifications []*UserNotificationDto) error {
	return nil
}
