package publishing

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Publish(notificationName string, message string, severity NotificationSeverity, userIds []int64, excludedUserIds []int64) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "publish",
			"notificationName", notificationName,
			"message", message,
			"severity", severity.String(),
			"userIds", fmt.Sprintf("%v", userIds),
			"excludedUserIds", fmt.Sprintf("%v", excludedUserIds),
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return s.Service.Publish(notificationName, message, severity, userIds, excludedUserIds)
}
