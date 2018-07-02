package publishing

import "errors"

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides publish methods.
type Service interface {
	Publish(
		notificationName string,
		message string,
		severity NotificationSeverity,
		userIds []int64,
		excludedUserIds []int64,
	) error
}

// NewService creates a publish service with necessary dependencies.
func NewService() Service {
	return &service{}
}

type service struct{}

// TODO ...
func (s *service) Publish(notificationName string, message string, severity NotificationSeverity, userIds []int64, excludedUserIds []int64) error {
	return nil
}
