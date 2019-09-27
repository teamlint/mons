package grpc

import (
	"context"
	"errors"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/teamlint/gox/convert"
	"github.com/teamlint/mons/sample/application/command"
	"github.com/teamlint/mons/sample/application/dto"
	endpoint "github.com/teamlint/mons/sample/application/endpoint"
	"github.com/teamlint/mons/sample/application/query"
	service "github.com/teamlint/mons/sample/application/service"
	pb "github.com/teamlint/mons/sample/application/transport/grpc/pb"
	grpc "google.golang.org/grpc"
)

// New returns a service backed by a gRPC server at the other end of the conn.
func New(conn *grpc.ClientConn, options map[string][]kitgrpc.ClientOption) (service.UserService, error) {
	findUserEndpoint := kitgrpc.NewClient(conn,
		"pb.User",
		"Find",
		encodeFindUserRequest,
		decodeFindUserResponse,
		pb.FindUserReply{},
		options["FindUser"]...,
	).Endpoint()
	updateUserEndpoint := kitgrpc.NewClient(
		conn,
		"pb.User",
		"Update",
		encodeUpdateUserRequest,
		decodeUpdateUserResponse,
		pb.UpdateUserReply{},
		options["UpdateUser"]...,
	).Endpoint()

	return endpoint.Endpoints{
		FindUserEndpoint:   findUserEndpoint,
		UpdateUserEndpoint: updateUserEndpoint,
	}, nil
}

// transport/grpc.EncodeRequestFunc that converts a user-domain request to a gRPC request.
func encodeFindUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*query.UserQuery)
	return &pb.FindUserRequest{Id: req.ID}, nil
}

// transport/grpc.DecodeResponseFunc that converts a gRPC concat reply to a user-domain concat response.
func decodeFindUserResponse(_ context.Context, reply interface{}) (interface{}, error) {
	rpl := reply.(*pb.FindUserReply)
	return &dto.User{
		ID:         rpl.Id,
		Username:   rpl.Username,
		Nickname:   rpl.Nickname,
		Intro:      rpl.Intro,
		IsApproved: rpl.IsApproved,
		CreatedAt:  convert.ToTime(rpl.CreatedAt),
		UpdatedAt:  convert.ToTime(rpl.UpdatedAt),
	}, nil
}

// transport/grpc.EncodeRequestFunc that converts a user-domain request to a gRPC request.
func encodeUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*command.UpdateUserCommand)
	return &pb.UpdateUserRequest{
		Id:       req.ID,
		Username: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
		Intro:    req.Intro,
	}, nil
}

// transport/grpc.DecodeResponseFunc that converts a gRPC concat reply to a user-domain concat response.
func decodeUpdateUserResponse(_ context.Context, reply interface{}) (interface{}, error) {
	if reply == nil {
		return nil, nil
	}
	rpl := reply.(*pb.UpdateUserReply)
	var err error
	if rpl.Err != "" {
		err = errors.New(rpl.Err)
	}
	return nil, err
}
