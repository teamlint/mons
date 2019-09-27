package endpoint

import (
	"context"

	kit "github.com/go-kit/kit/endpoint"
    api "github.com/teamlint/mons/sample2/services/user/api"
)

// Endpoints User endpoints
type Endpoints struct {
    FindEndpoint kit.Endpoint
    UpdateEndpoint kit.Endpoint
}

func New(svc api.UserServer, mdw map[string][]kit.Middleware) Endpoints {
	eps := Endpoints{
		FindEndpoint:   MakeFindEndpoint(svc),
		UpdateEndpoint:   MakeUpdateEndpoint(svc),
	}
	for _, m := range mdw["Find"] {
		eps.FindEndpoint = m(eps.FindEndpoint)
	}
	for _, m := range mdw["Update"] {
		eps.UpdateEndpoint = m(eps.UpdateEndpoint)
	}
	return eps
}

func (e Endpoints) Find(ctx context.Context, in *api.FindUserRequest) (*api.FindUserReply, error) {
    out, err := e.FindEndpoint(ctx, in)
    if err !=nil {
        return nil, err
    }
    return out.(*api.FindUserReply), nil
}

func (e Endpoints) Update(ctx context.Context, in *api.UpdateUserRequest) (*api.UpdateUserReply, error) {
    out, err := e.UpdateEndpoint(ctx, in)
    if err !=nil {
        return nil, err
    }
    return out.(*api.UpdateUserReply), nil
}



func MakeFindEndpoint(svc api.UserServer) kit.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(*api.FindUserRequest)
        resp, err := svc.Find(ctx, req)
        return resp, err
    }
}

func MakeUpdateEndpoint(svc api.UserServer) kit.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(*api.UpdateUserRequest)
        resp, err := svc.Update(ctx, req)
        return resp, err
    }
}

