package nats

import (
	"context"
	"encoding/json"

    "github.com/teamlint/mons/sample2/services/user/api"
    "github.com/teamlint/mons/sample2/services/user/internal/endpoint"

	kitnats "github.com/go-kit/kit/transport/nats"
	"github.com/nats-io/nats.go"
)

type natsServer struct {
	FindHandler   *kitnats.Subscriber
	UpdateHandler   *kitnats.Subscriber
}

// NewNATSServer
func NewNATSServer(endpoints endpoint.Endpoints, options map[string][]kitnats.SubscriberOption) natsServer {
	return natsServer{
		FindHandler:   makeFindHandler(endpoints, options["Find"]),
		UpdateHandler:   makeUpdateHandler(endpoints, options["Update"]),
	}
}

func makeFindHandler(endpoints endpoint.Endpoints, options []kitnats.SubscriberOption) *kitnats.Subscriber {
	handler := kitnats.NewSubscriber(
		endpoints.FindEndpoint,
		decodeFindRequest,
		kitnats.EncodeJSONResponse,
		options...,
	)
	return handler
}

func makeUpdateHandler(endpoints endpoint.Endpoints, options []kitnats.SubscriberOption) *kitnats.Subscriber {
	handler := kitnats.NewSubscriber(
		endpoints.UpdateEndpoint,
		decodeUpdateRequest,
		kitnats.EncodeJSONResponse,
		options...,
	)
	return handler
}

func decodeFindRequest(ctx context.Context, msg *nats.Msg) (interface{}, error) {
	var request api.FindUserRequest
	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}
	return &request, nil
}

func decodeUpdateRequest(ctx context.Context, msg *nats.Msg) (interface{}, error) {
	var request api.UpdateUserRequest
	if err := json.Unmarshal(msg.Data, &request); err != nil {
		return nil, err
	}
	return &request, nil
}

