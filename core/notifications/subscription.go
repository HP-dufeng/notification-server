package notifications

// SubscriptionInfo is the central class in the domain model.
type SubscriptionInfo struct {
	ID               ID
	NotificationName string
	UserID           int64
}

// SubscriptionInfoRepository provides access a subscription info store.
type SubscriptionInfoRepository interface {
	Insert(s *SubscriptionInfo) error
	GetAll(notificationName string) []SubscriptionInfo
}
