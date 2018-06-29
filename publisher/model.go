package publisher

// NotificationData used to store data for a notification.
type NotificationData struct {
	Properties map[string]interface{}
}

// NotificationSeverity notification severity.
type NotificationSeverity int8

// Notification severity.
const (
	Info NotificationSeverity = iota
	Success
	Warn
	Error
	Fatal
)

var severityStrings = map[string]NotificationSeverity{
	"info":    Info,
	"success": Success,
	"warn":    Warn,
	"error":   Error,
	"fatal":   Fatal,
}

func (n NotificationSeverity) String() string {
	for k, v := range severityStrings {
		if n == v {
			return k
		}
	}

	return ""
}

// ParseSeverity pase string to NotificationSeverity type
func ParseSeverity(s string) NotificationSeverity {
	return severityStrings[s]
}
