package notifications

import "time"

// UserNotificationDto represents a notification sent to a user.
type UserNotificationDto struct {
	UserID                int64
	UserNotificationState string
	NotificationName      string
	Severity              string
	Data                  string
}

// UserNotificationInfo is the central class in the domain model.
type UserNotificationInfo struct {
	ID             ID
	UserID         int64
	NotificationID ID
	State          UserNotificationState
	CreationTime   int64
}

// NewUserNotificationInfo init a new instance of UserNotificationInfo
func NewUserNotificationInfo(userID int64, notificationID ID) *UserNotificationInfo {
	return &UserNotificationInfo{
		ID:             NextID(),
		UserID:         userID,
		NotificationID: notificationID,
		State:          Unread,
		CreationTime:   time.Now().Unix(),
	}
}

// UserNotificationInfoRepository provides access a user_notification store.
type UserNotificationInfoRepository interface {
	Insert(userNotificationInfo *UserNotificationInfo) error
	FindAll() []*UserNotificationInfo
}

// UserNotificationState describes state of notification.
type UserNotificationState int8

// Valid user notification state.
const (
	Unread UserNotificationState = iota
	Read
)

var userNotificationStateStrings = map[string]UserNotificationState{
	"unread": Unread,
	"read":   Read,
}

func (s UserNotificationState) String() string {
	for k, v := range userNotificationStateStrings {
		if v == s {
			return k
		}
	}

	return ""
}

// ParseState pase string to Severity type
func ParseState(s string) UserNotificationState {
	return userNotificationStateStrings[s]
}
