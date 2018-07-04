package notifications

import (
	"errors"
)

// NotificationInfo is the central class in the domain model.
type NotificationInfo struct {
	ID               ID
	NotificationName string
	Severity         Severity
	Data             string
}

// NotificationInfoRepository provides access a notification store.
type NotificationInfoRepository interface {
	Insert(notification *NotificationInfo) error
	Find(id ID) (*NotificationInfo, error)
}

// ErrUnknown is used when a notification could not be found.
var ErrUnknown = errors.New("unknown notification")

// Severity describes severity of notification.
type Severity int8

// Valid notification severity.
const (
	Info Severity = iota
	Success
	Warn
	Error
	Fatal
)

var severityStrings = map[string]Severity{
	"info":    Info,
	"success": Success,
	"warn":    Warn,
	"error":   Error,
	"fatal":   Fatal,
}

func (n Severity) String() string {
	for k, v := range severityStrings {
		if n == v {
			return k
		}
	}

	return ""
}

// ParseSeverity pase string to Severity type
func ParseSeverity(s string) Severity {
	return severityStrings[s]
}
