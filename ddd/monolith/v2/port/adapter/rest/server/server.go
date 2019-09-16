package server

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/teamlint/iris"
	"github.com/teamlint/mons/ddd/monolith/v2/application/service"
	doservice "github.com/teamlint/mons/ddd/monolith/v2/domain/model/account/service"
	"github.com/teamlint/mons/ddd/monolith/v2/port/adapter/event"
	"github.com/teamlint/mons/ddd/monolith/v2/port/adapter/persistence/repository"
	"github.com/teamlint/mons/ddd/monolith/v2/port/adapter/persistence/repository/mysql"
	"github.com/teamlint/mons/ddd/monolith/v2/port/adapter/rest/handler"
	serviceimpl "github.com/teamlint/mons/ddd/monolith/v2/port/adapter/service"
	"go.uber.org/fx"
)

var params service.UserServiceParams

type Server struct {
	*iris.Application
	*iris.Configuration
	DB *gorm.DB
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			NewDB, // DB
			repository.NewGormRepositoryContext, // repository context
			mysql.NewUserRepository,             //  repository
			doservice.NewUserService,            // domain service
			service.NewUserService,              // application service
			serviceimpl.NewLogService,           // log service
			handler.NewUserHandler,              // adapter rest handler
			NewConfig,                           // config
			// event.New,                           // event
			// event.NewEhEvent, // event
			event.NewNatsEvent, // event
			New,                // web server
		),
		// fx.Provide(fx.Annotated{Name: "event", Target: event.NewEhEvent}), // event,another inject method
		// fx.Provide(fx.Annotated{Name: "foo", Source: new(doevent.Event), Target: event.NewEhEvent}), // event,another inject method
		fx.Populate(&params),
		fx.Invoke(Bootstrap), // bootstrap
	)
}
func New(cfg *iris.Configuration, db *gorm.DB, userHandler *handler.UserHandler) *Server {
	// app
	server := Server{Configuration: cfg, DB: db}
	server.Application = iris.New()
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
			log.Println("Golang DDD practice v2.3")
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
func Run() {
	var g fx.DotGraph
	app := fx.New(fx.Populate(&g), Module())
	// log.Printf("app dependency graph: %v\n", g)
	log.Printf("user service params: %v\n", params)
	app.Run()
}

// NewDB database
func NewDB(cfg *iris.Configuration) *gorm.DB {
	// init db
	// connStr := "root:123456@tcp(localhost:3306)/mons?charset=utf8mb4&parseTime=True&loc=Local"
	connStr := cfg.GetOther()["DBConnectionString"]
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(3 * time.Minute)
	return db
}

// NewConfig init config
func NewConfig() *iris.Configuration {
	cfg := iris.YAML("./config.yml")
	return &cfg
}
