package nats

import (
	"context"
	"encoding/json"

	kitnats "github.com/go-kit/kit/transport/nats"
	"github.com/nats-io/nats.go"
	api "github.com/teamlint/mons/sample2/services/user/api"
)

func ErrorEncoder(ctx context.Context, err error, reply string, nc *nats.Conn) {
	var response api.ErrorResponse
	if err != nil {
		response.Error = &api.Error{Code: "_internal_", Message: err.Error()}
	}
	b, err := json.Marshal(response)
	if err != nil {
		return
	}
	if err := nc.Publish(reply, b); err != nil {
		response.Error = &api.Error{Code: "_internal_", Message: err.Error()}
	}
}
func SubscriberErrorOption() kitnats.SubscriberOption {
	return kitnats.SubscriberErrorEncoder(ErrorEncoder)
}
