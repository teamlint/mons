package server

import (
	"context"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/teamlint/iris"
	"github.com/teamlint/mons/ddd/monolith/facade/rest"
	"go.uber.org/fx"
)

type Server struct {
	*iris.Application
	*iris.Configuration
	DB *gorm.DB
}

func New(cfg *iris.Configuration, db *gorm.DB, userHandler *rest.UserHandler) *Server {
	// app
	server := Server{Configuration: cfg, DB: db}
	server.Application = iris.Default()
	server.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello DDD")
	})
	server.Get("/user/{name}", userHandler.GetUser)
	return &server
}
func Bootstrap(life fx.Lifecycle, server *Server) {
	// life cycle
	life.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 30 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			log.Println("Golang DDD practice v1.0")
			log.Println("Starting HTTP server.")
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			// go server.ListenAndServe()
			// server
			addr := fmt.Sprintf(":%v", server.Configuration.GetOther()["ServerPort"])
			go server.Run(iris.Addr(addr))
			// server.Run(iris.Addr(addr))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping HTTP server.")
			// log.Println("[DBStatus]", server.DB.DB().Stats())
			server.DB.DB().Close()
			log.Println("Stopping DB server.")
			return server.Shutdown(ctx)
		},
	})
}
