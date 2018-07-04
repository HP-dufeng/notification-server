package managing

import (
	"errors"

	"github.com/fengdu/notification-server/core/notifications"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides user notifications manage methods.
type Service interface {
	GetUserNotifications(userID int64) []*notifications.UserNotificationDto
}

// NewService creates a user notification manage service with necessary dependencies.
func NewService(userNotificationManager notifications.UserNotificationManager) Service {
	return &service{
		userNotificationManager,
	}
}

type service struct {
	userNotificationManager notifications.UserNotificationManager
}

func (s *service) GetUserNotifications(userID int64) []*notifications.UserNotificationDto {
	return s.userNotificationManager.GetUserNotifications(userID)
}
