package publishing

import (
	"errors"

	"github.com/fengdu/notification-server/core/notifications"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides publish methods.
type Service interface {
	Publish(
		notificationName string,
		message string,
		severity notifications.Severity,
		userIds []int64,
		excludedUserIds []int64,
	) error
}

// NewService creates a publish service with necessary dependencies.
func NewService(notificationPublisher notifications.Publisher) Service {
	return &service{
		notificationPublisher,
	}
}

type service struct {
	notificationPublisher notifications.Publisher
}

// TODO ...
func (s *service) Publish(notificationName string, message string, severity notifications.Severity, userIds []int64, excludedUserIds []int64) error {

	err := s.notificationPublisher.Publish(notificationName, message, severity, userIds, excludedUserIds)

	return err
}
