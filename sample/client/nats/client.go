package nats

import (
	"context"
	"encoding/json"

	"github.com/teamlint/mons/sample/application/dto"
	"github.com/teamlint/mons/sample/application/endpoint"
	"github.com/teamlint/mons/sample/application/service"

	kitnats "github.com/go-kit/kit/transport/nats"

	"github.com/nats-io/nats.go"
)

// func New(nc *nats.Conn) service.StringService {
// 	return natsserver.NewNATSClient(nc)
// }

// client
// func NewNATSClient(nc *nats.Conn) service.StringService {
func New(nc *nats.Conn, options map[string][]kitnats.PublisherOption) (service.UserService, error) {
	finduser := kitnats.NewPublisher(
		nc,
		"user.finduser",
		kitnats.EncodeJSONRequest,
		decodeFindUserResponse,
		options["FindUser"]...,
	).Endpoint()

	updateuser := kitnats.NewPublisher(
		nc,
		"user.updateuser",
		kitnats.EncodeJSONRequest,
		decodeUpdateUserResponse,
		options["UpdateUser"]...,
	).Endpoint()

	return endpoint.Endpoints{
		FindUserEndpoint:   finduser,
		UpdateUserEndpoint: updateuser,
	}, nil
}

func decodeFindUserResponse(_ context.Context, msg *nats.Msg) (interface{}, error) {
	// var response endpoint.UppercaseResponse
	var response dto.User

	if err := json.Unmarshal(msg.Data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func decodeUpdateUserResponse(_ context.Context, msg *nats.Msg) (interface{}, error) {
	// var response domain.User
	var response interface{}

	if err := json.Unmarshal(msg.Data, &response); err != nil {
		return nil, err
	}

	return response, nil
	// return nil, nil
}
