package notifications

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
