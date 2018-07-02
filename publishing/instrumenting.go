package publishing

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) Publish(notificationName string, message string, severity NotificationSeverity, userIds []int64, excludedUserIds []int64) error {
	defer func(begin time.Time) {
		s.requestCount.With("method", "publish").Add(1)
		s.requestLatency.With("method", "publish").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Publish(notificationName, message, severity, userIds, excludedUserIds)
}
