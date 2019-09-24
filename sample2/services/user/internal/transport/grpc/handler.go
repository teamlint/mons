package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
    api "github.com/teamlint/mons/sample2/services/user/api"
    "github.com/teamlint/mons/sample2/services/user/internal/endpoint"
)

type grpcServer struct {
	FindHandler   kitgrpc.Handler
	UpdateHandler   kitgrpc.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC server
func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]kitgrpc.ServerOption) api.UserServer {
	return &grpcServer{
		FindHandler:   makeFindHandler(endpoints, options["Find"]),
		UpdateHandler:   makeUpdateHandler(endpoints, options["Update"]),
	}
}

func makeFindHandler(endpoints endpoint.Endpoints, options []kitgrpc.ServerOption) kitgrpc.Handler {
	return kitgrpc.NewServer(endpoints.FindEndpoint, decodeRequest, encodeFindResponse, options...)
}

func makeUpdateHandler(endpoints endpoint.Endpoints, options []kitgrpc.ServerOption) kitgrpc.Handler {
	return kitgrpc.NewServer(endpoints.UpdateEndpoint, decodeRequest, encodeUpdateResponse, options...)
}

func decodeRequest(ctx context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeFindResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*api.FindUserReply)
	return resp, nil
}

func encodeUpdateResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*api.UpdateUserReply)
	return resp, nil
}

func (g *grpcServer) Find(ctx context.Context, req *api.FindUserRequest) (*api.FindUserReply, error) {
	_, resp, err := g.FindHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*api.FindUserReply), nil
}

func (g *grpcServer) Update(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserReply, error) {
	_, resp, err := g.UpdateHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*api.UpdateUserReply), nil
}

