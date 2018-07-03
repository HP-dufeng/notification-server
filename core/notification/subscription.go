package notification

// SubscriptionInfo is the central class in the domain model.
type SubscriptionInfo struct {
	ID               ID
	NotificationName string
	UserID           int64
}

// SubscriptionRepository provides access a subscription info store.
type SubscriptionRepository interface {
	Insert(s *SubscriptionInfo) error
	GetAll(notificationName string) []SubscriptionInfo
}
