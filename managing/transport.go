package managing

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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

	getUserNotificationsHandler := kithttp.NewServer(
		makeGetUserNotificationsEndpoint(s),
		decodeGetUserNotificationsRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/managing/v1/user/{id}", getUserNotificationsHandler).Methods("GET")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetUserNotificationsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, ErrInvalidArgument
	}

	return getUserNotificationsRequest{
		userID: userID,
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
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
