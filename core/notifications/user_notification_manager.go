package notifications

// UserNotificationManager used to manage user notifications.
type UserNotificationManager interface {
	GetUserNotifications(userID int64) []*UserNotificationDto
}

type userNotificationManager struct {
	store Store
}

// NewUserNotificationManager returns a new instance of a  UserNotificationManager interface.
func NewUserNotificationManager(s Store) UserNotificationManager {
	return &userNotificationManager{
		store: s,
	}
}

func (u *userNotificationManager) GetUserNotifications(userID int64) []*UserNotificationDto {
	userNotifications := u.store.GetUserNotifications(userID)
	dtos := make([]*UserNotificationDto, 0, len(userNotifications))
	for _, v := range userNotifications {
		if n, err := u.store.GetNotification(v.NotificationID); err == nil {
			dtos = append(dtos, &UserNotificationDto{
				UserID:                v.UserID,
				UserNotificationState: v.State.String(),
				NotificationName:      n.NotificationName,
				Severity:              n.Severity.String(),
				Data:                  n.Data,
			})
		}
	}

	return dtos
}
