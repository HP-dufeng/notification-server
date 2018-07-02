package publishing

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the publisher service.
func MakeHandler(s Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	publishHandler := kithttp.NewServer(
		makePublishEndpoint(s),
		decodePublishRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/publish", publishHandler).Methods("POST")

	return r
}

func decodePublishRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		NotificationName string  `json:"notificationName"`
		Message          string  `json:"message"`
		Severity         string  `json:"severity"`
		UserIds          []int64 `json:"userIds"`
		ExcludedUserIds  []int64 `json:"excludedUserIds"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return publishRequest{
		notificationName: body.NotificationName,
		message:          body.Message,
		severity:         ParseSeverity(body.Severity),
		userIds:          body.UserIds,
		excludedUserIds:  body.ExcludedUserIds,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	// case cargo.ErrUnknown:
	// 	w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
