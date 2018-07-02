package publishing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type publishRequest struct {
	notificationName string
	message          string
	severity         NotificationSeverity
	userIds          []int64
	excludedUserIds  []int64
}

type publishResponse struct {
	Err error `json:"error,omitempty"`
}

func (p publishResponse) error() error { return p.Err }

func makePublishEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(publishRequest)
		err := s.Publish(req.notificationName, req.message, req.severity, req.userIds, req.excludedUserIds)

		return publishResponse{Err: err}, nil
	}
}
