package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/teamlint/gox/timex"
	"github.com/teamlint/mons/sample/application/command"
	"github.com/teamlint/mons/sample/application/dto"
	endpoint "github.com/teamlint/mons/sample/application/endpoint"
	"github.com/teamlint/mons/sample/application/query"
	pb "github.com/teamlint/mons/sample/application/transport/grpc/pb"
)

// NewGRPCServer makes a set of endpoints available as a gRPC server
type grpcServer struct {
	findUser   kitgrpc.Handler
	updateUser kitgrpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]kitgrpc.ServerOption) pb.UserServer {
	return &grpcServer{
		findUser:   makeFindUserHandler(endpoints, options["FindUser"]),
		updateUser: makeUpdateUserHandler(endpoints, options["UpdateUser"]),
	}
}

func makeFindUserHandler(endpoints endpoint.Endpoints, options []kitgrpc.ServerOption) kitgrpc.Handler {
	return kitgrpc.NewServer(endpoints.FindUserEndpoint, decodeFindUserRequest, encodeFindUserResponse, options...)
}

// transport/grpc.DecodeRequestFunc that converts a gRPC request to a user-domain request.
func decodeFindUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.FindUserRequest)
	return &query.UserQuery{ID: req.Id}, nil
}

// transport/grpc.EncodeResponseFunc that converts a user-domain response to a gRPC reply.
func encodeFindUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	if r == nil {
		return nil, nil
	}
	resp := r.(*dto.User)
	reply := pb.FindUserReply{
		Id:         resp.ID,
		Username:   resp.Username,
		Nickname:   resp.Nickname,
		Intro:      resp.Intro,
		Password:   resp.Password,
		IsApproved: resp.IsApproved,
		CreatedAt:  timex.Format(resp.CreatedAt),
		UpdatedAt:  timex.Format(resp.UpdatedAt),
	}
	return &reply, nil
}
func (g *grpcServer) Find(ctx context.Context, req *pb.FindUserRequest) (*pb.FindUserReply, error) {
	_, rep, err := g.findUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.FindUserReply), nil
}

func makeUpdateUserHandler(endpoints endpoint.Endpoints, options []kitgrpc.ServerOption) kitgrpc.Handler {
	return kitgrpc.NewServer(endpoints.UpdateUserEndpoint, decodeUpdateUserRequest, encodeUpdateUserResponse, options...)
}

// transport/grpc.DecodeRequestFunc that converts a gRPC request to a user-domain request.
func decodeUpdateUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.UpdateUserRequest)
	return &command.UpdateUserCommand{
		ID:       req.Id,
		Username: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
		Intro:    req.Intro,
	}, nil
}

// transport/grpc.EncodeResponseFunc that converts a user-domain response to a gRPC reply.
func encodeUpdateUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	if r == nil {
		return &pb.UpdateUserReply{}, nil
	}
	resp := r.(error)
	var reply pb.UpdateUserReply
	if resp != nil {
		reply.Err = resp.Error()
	}
	return &reply, resp
}
func (g *grpcServer) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	_, rep, err := g.updateUser.ServeGRPC(ctx, req)
	if rep == nil && err == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserReply), nil
}
