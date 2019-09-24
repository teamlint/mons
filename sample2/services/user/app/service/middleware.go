package service

import (
	"context"

    "github.com/teamlint/mons/sample2/services/user/api"

	kitlog "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(api.UserServer) api.UserServer

type loggingMiddleware struct {
	logger kitlog.Logger
	next   api.UserServer
}

// LoggingMiddleware takes a logger as a dependency
// and returns a api.UserServer Middleware.
func LoggingMiddleware(logger kitlog.Logger) Middleware {
	return func(next api.UserServer) api.UserServer {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) Find(ctx context.Context, req *api.FindUserRequest) (reply *api.FindUserReply, err error) {
	defer func() {
		_ = l.logger.Log("method", "Find", "request", req, "reply", reply, "err", err)
	}()
	return l.next.Find(ctx, req)
}

func (l loggingMiddleware) Update(ctx context.Context, req *api.UpdateUserRequest) (reply *api.UpdateUserReply, err error) {
	defer func() {
		_ = l.logger.Log("method", "Update", "request", req, "reply", reply, "err", err)
	}()
	return l.next.Update(ctx, req)
}

