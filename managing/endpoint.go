package managing

import (
	"context"

	"github.com/fengdu/notification-server/core/notifications"
	"github.com/go-kit/kit/endpoint"
)

type getUserNotificationsRequest struct {
	userID int64
}

type getUserNotificationshResponse struct {
	Datas []*notifications.UserNotificationDto
	Err   error `json:"error,omitempty"`
}

func (p getUserNotificationshResponse) error() error { return p.Err }

func makeGetUserNotificationsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserNotificationsRequest)
		datas := s.GetUserNotifications(req.userID)

		return getUserNotificationshResponse{Datas: datas}, nil
	}
}
