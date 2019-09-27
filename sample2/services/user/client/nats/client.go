package nats

import (
	"context"
	"encoding/json"
	"errors"

    api "github.com/teamlint/mons/sample2/services/user/api"
    "github.com/teamlint/mons/sample2/services/user/internal/endpoint"

	kitnats "github.com/go-kit/kit/transport/nats"

	"github.com/nats-io/nats.go"
)

// NewNATSClient
func New(nc *nats.Conn, options map[string][]kitnats.PublisherOption) (api.UserServer, error) {
	findEndpoint := kitnats.NewPublisher(
		nc,
		"User.Find",
		kitnats.EncodeJSONRequest,
		decodeFindResponse,
		options["Find"]...,
	).Endpoint()
    
	updateEndpoint := kitnats.NewPublisher(
		nc,
		"User.Update",
		kitnats.EncodeJSONRequest,
		decodeUpdateResponse,
		options["Update"]...,
	).Endpoint()
    
	return endpoint.Endpoints{
		FindEndpoint:   findEndpoint,
		UpdateEndpoint:   updateEndpoint,
	}, nil
}

func decodeFindResponse(ctx context.Context, msg *nats.Msg) (interface{}, error) {
	var errResponse api.ErrorResponse
	if err := json.Unmarshal(msg.Data, &errResponse); err != nil {
		return nil, err
	}
	if errResponse.Error != nil && errResponse.Error.Code == "_internal_" {
		return nil, errors.New(errResponse.Error.Message)
	}
	var response api.FindUserReply
	if err := json.Unmarshal(msg.Data, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func decodeUpdateResponse(ctx context.Context, msg *nats.Msg) (interface{}, error) {
	var errResponse api.ErrorResponse
	if err := json.Unmarshal(msg.Data, &errResponse); err != nil {
		return nil, err
	}
	if errResponse.Error != nil && errResponse.Error.Code == "_internal_" {
		return nil, errors.New(errResponse.Error.Message)
	}
	var response api.UpdateUserReply
	if err := json.Unmarshal(msg.Data, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

