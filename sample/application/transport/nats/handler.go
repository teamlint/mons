package nats

import (
	"context"
	"encoding/json"

	"github.com/teamlint/mons/sample/application/command"
	"github.com/teamlint/mons/sample/application/endpoint"
	"github.com/teamlint/mons/sample/application/query"

	kitnats "github.com/go-kit/kit/transport/nats"
	"github.com/nats-io/nats.go"
)

type natsServer struct {
	FindUser   *kitnats.Subscriber
	UpdateUser *kitnats.Subscriber
}

// NewNATSServer
func NewNATSServer(endpoints endpoint.Endpoints, options map[string][]kitnats.SubscriberOption) natsServer {
	return natsServer{
		FindUser:   makeFindUserHandler(endpoints, options["FindUser"]),
		UpdateUser: makeUpdateUserHandler(endpoints, options["UpdateUser"]),
	}
}
func makeFindUserHandler(endpoints endpoint.Endpoints, options []kitnats.SubscriberOption) *kitnats.Subscriber {
	handler := kitnats.NewSubscriber(
		endpoints.FindUserEndpoint,
		decodeFindUserRequest,
		kitnats.EncodeJSONResponse,
		options...,
	)
	return handler
}
func makeUpdateUserHandler(endpoints endpoint.Endpoints, options []kitnats.SubscriberOption) *kitnats.Subscriber {
	handler := kitnats.NewSubscriber(
		endpoints.UpdateUserEndpoint,
		decodeUpdateUserRequest,
		kitnats.EncodeJSONResponse,
		options...,
	)
	return handler
}
func decodeFindUserRequest(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var request query.UserQuery

	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}
	return &request, nil
}

func decodeUpdateUserRequest(_ context.Context, msg *nats.Msg) (interface{}, error) {
	var request command.UpdateUserCommand

	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}
	return &request, nil
}
