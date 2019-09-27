package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
    "github.com/teamlint/mons/sample2/services/user/internal/endpoint"
    api "github.com/teamlint/mons/sample2/services/user/api"
	grpc "google.golang.org/grpc"
)

// New returns a service backed by a gRPC server at the other end of the conn.
func New(conn *grpc.ClientConn, options map[string][]kitgrpc.ClientOption) (api.UserServer, error) {
	findEndpoint := kitgrpc.NewClient(
        conn,
		"api.User",
		"Find",
		encodeFindRequest,
		decodeFindResponse,
		api.FindUserReply{},
		options["Find"]...,
	).Endpoint()
	updateEndpoint := kitgrpc.NewClient(
        conn,
		"api.User",
		"Update",
		encodeUpdateRequest,
		decodeUpdateResponse,
		api.UpdateUserReply{},
		options["Update"]...,
	).Endpoint()
	return endpoint.Endpoints{
		FindEndpoint:   findEndpoint,
		UpdateEndpoint:   updateEndpoint,
	}, nil
}
func encodeFindRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*api.FindUserRequest)
	return req, nil
}

func decodeFindResponse(ctx context.Context, reply interface{}) (interface{}, error) {
	rpl := reply.(*api.FindUserReply)
    return rpl, nil
}
func encodeUpdateRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*api.UpdateUserRequest)
	return req, nil
}

func decodeUpdateResponse(ctx context.Context, reply interface{}) (interface{}, error) {
	rpl := reply.(*api.UpdateUserReply)
    return rpl, nil
}

