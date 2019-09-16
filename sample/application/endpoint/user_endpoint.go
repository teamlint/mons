package endpoint

import (
	"context"
	"errors"
	"log"

	kit "github.com/go-kit/kit/endpoint"
	"github.com/teamlint/mons/sample/application/command"
	"github.com/teamlint/mons/sample/application/dto"
	"github.com/teamlint/mons/sample/application/query"
	"github.com/teamlint/mons/sample/application/service"
)

// Endpoints
type Endpoints struct {
	FindUserEndpoint   kit.Endpoint
	UpdateUserEndpoint kit.Endpoint
}

func New(s service.UserService, mdw map[string][]kit.Middleware) Endpoints {
	eps := Endpoints{
		FindUserEndpoint:   MakeFindUserEndpoint(s),
		UpdateUserEndpoint: MakeUpdateUserEndpoint(s),
	}
	for _, m := range mdw["finduser"] {
		eps.FindUserEndpoint = m(eps.FindUserEndpoint)
	}
	for _, m := range mdw["updateuser"] {
		eps.UpdateUserEndpoint = m(eps.UpdateUserEndpoint)
	}
	return eps
}
func (e Endpoints) String() string {
	return "[FindUserEndpoint,UpdateUserEndpoint]"
}

func (e Endpoints) Find(ctx context.Context, query *query.UserQuery) (*dto.User, error) {
	resp, err := e.FindUserEndpoint(ctx, query)
	if err != nil {
		return nil, err
	}
	// log.Printf("[Endpoint]  FindUser: <<%v,%T>>, err: <<%v>>\n", resp, resp, err)
	r, ok := resp.(*dto.User)
	if ok {
		return r, nil
	}
	return nil, errors.New("类型转换失败")

}
func (e Endpoints) Update(ctx context.Context, cmd *command.UpdateUserCommand) error {
	_, err := e.UpdateUserEndpoint(ctx, cmd)
	return err
}

func MakeFindUserEndpoint(svc service.UserService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*query.UserQuery)
		resp, err := svc.Find(ctx, req)
		log.Printf("[Endpoint] resp: <<%v,%T>>, err: <<%v>>\n", resp, resp, err)
		return resp, err
	}
}
func MakeUpdateUserEndpoint(svc service.UserService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*command.UpdateUserCommand)
		err := svc.Update(ctx, req)
		// log.Printf("[Endpoint] resp: <<%v,%T>>, err: <<%v>>\n", resp, resp, err)
		return nil, err
	}
}
