package server

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/teamlint/mons/container"
	"github.com/teamlint/mons/sample/adapter/repository/mysql"
	"github.com/teamlint/mons/sample/application/endpoint"
	"github.com/teamlint/mons/sample/application/service"
	grpchandler "github.com/teamlint/mons/sample/application/transport/grpc"
	pb "github.com/teamlint/mons/sample/application/transport/grpc/pb"
	httphandler "github.com/teamlint/mons/sample/application/transport/http"
	natshandler "github.com/teamlint/mons/sample/application/transport/nats"
	"github.com/teamlint/run"
	"google.golang.org/grpc"

	// natshttp "github.com/teamlint/mons/sample/adapter/transport/natshttp"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	kitnats "github.com/go-kit/kit/transport/nats"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/nats-io/nats.go"
	"github.com/prometheus/client_golang/prometheus"
	svcimpl "github.com/teamlint/mons/sample/adapter/application/service"
	dosvcimpl "github.com/teamlint/mons/sample/adapter/domain/service"

	"github.com/defval/inject"
	"github.com/teamlint/mons/sample/application/event"
	dosvc "github.com/teamlint/mons/sample/domain/service"
	shared "github.com/teamlint/mons/shared/adapter/repository"
	sharedevent "github.com/teamlint/mons/shared/application/event"
	natsevent "github.com/teamlint/mons/shared/application/event/nats"
)

var (
	fs       = flag.NewFlagSet("arch", flag.ExitOnError)
	connStr  = "root:123456@tcp(localhost:3306)/mons?charset=utf8mb4&parseTime=True&loc=Local"
	dbConn   = fs.String("db-conn", connStr, "URL for connection to database")
	natsAddr = fs.String("nats-addr", nats.DefaultURL, "NATS listen address")
	httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
	//  natsHTTP = fs.String("nats-http", ":8082", "NATS over HTTP listen address")
	grpcAddr = fs.String("grpc-addr", ":8082", "GRPC listen address")
	mode     = fs.String("mode", "normal", "run mode [normal,di,ij]")
	debug    = fs.Bool("debug", true, "debug mode")
	logger   kitlog.Logger
)

func Run() {
	// args
	err := fs.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	switch *mode {
	case "di":
		di()
	case "ij":
		ij()
	default:
		normal()
	}
}

func normal() {
	// logger
	logger = NewLogger()
	// db
	db, err := NewDB()
	if err != nil {
		logger.Log(err)
	}
	defer db.Close()
	// nats
	nc, err := NewNATS()
	if err != nil {
		logger.Log(err)
	}
	defer nc.Close()
	// repository context
	repoCtx := shared.NewGormRepositoryContext(db)
	// repository
	repo := mysql.NewUserRepository()
	// event config
	evtConfig := natsevent.Config{
		ClusterID: "test-cluster",
		ClientID:  "nats_client_mons_normal",
		URL:       *natsAddr,
	}
	// eventer
	eventer := natsevent.NewNATSEventer(&evtConfig)
	// domain user service
	doUserSvc := dosvcimpl.NewUserService()
	// service
	svcConf := svcimpl.UserServiceConfig{
		UserRepo:    repo,
		RepoContext: repoCtx,
		UserSvc:     doUserSvc,
		Eventer:     eventer,
	}
	userSvc := svcimpl.NewUserService(svcConf)
	userSvc = service.NewUserService(userSvc, []service.UserMiddleware{})
	// endpoint
	eps := endpoint.New(userSvc, getEndpointMiddleware(logger))
	// subscribe
	eventSub(eventer)
	// server
	g := NewServer(eps, nc)

	// interrupt
	initCancelInterrupt(g)
	log.Println("exit", g.Run())
}
func di() {
	c := container.New()
	// logger
	c.Provide(NewLogger)
	// db
	c.Provide(NewDB)
	// nats
	c.Provide(NewNATS)
	// repository
	c.Provide(shared.NewGormRepositoryContext)
	c.Provide(mysql.NewUserRepository)
	// event config
	c.Provide(func() *natsevent.Config {
		evtConfig := natsevent.Config{
			ClusterID: "test-cluster",
			ClientID:  "nats_client_mons_di",
			URL:       *natsAddr,
		}
		return &evtConfig
	})
	// eventer
	c.Provide(natsevent.NewNATSEventer)
	// service
	// c.Provide(dosvcimpl.NewUserService)
	c.Provide(func() dosvc.UserService {
		return nil
	})
	c.Provide(svcimpl.NewUserService)
	// c.Provide(service.NewUserService)
	// // middleware
	// c.Provide(NewMiddleware())
	// // endpoint
	// c.Provide(endpoint.New)
	// // server
	// c.Provide(NewServer)
	var nc *nats.Conn
	{
		err := c.Invoke(func(c *nats.Conn) {
			nc = c
		})
		if err != nil {
			panic(fmt.Sprintf("[NATS] error: %v", err))
		}
	}
	// logger
	{
		err := c.Invoke(func(l kitlog.Logger) {
			logger = l
		})
		if err != nil {
			panic(fmt.Sprintf("[Logger] error: %v", err))
		}
	}

	// subscribe
	{
		err := c.Invoke(func(eventer sharedevent.Eventer) {
			eventSub(eventer)
		})
		if err != nil {
			panic(fmt.Sprintf("[Eventer] error: %v", err))
		}
	}
	// run
	{
		err := c.Invoke(func(svc service.UserService) {
			userSvc := service.NewUserService(svc, []service.UserMiddleware{})
			eps := endpoint.New(userSvc, getEndpointMiddleware(logger))
			// log.Printf("[UserService] eps value: %+v", eps)
			g := NewServer(eps, nc)
			initCancelInterrupt(g)
			log.Println("exit", g.Run())
		})
		if err != nil {
			panic(fmt.Sprintf("[UserService] error: %v", err))
		}
	}
}

func ij() {
	file, err := os.Create("di.gv")
	if err != nil {
		panic(err)
	}
	userServiceBundle := inject.Bundle(
		// db
		inject.Provide(NewDB),
		// repository context
		inject.Provide(shared.NewGormRepositoryContext),
		// repository
		inject.Provide(mysql.NewUserRepository),
		// service config
		inject.Provide(svcimpl.NewUserServiceConfig),
		// service
		inject.Provide(svcimpl.NewUserService),
	)
	c, err := inject.New(
		// logger
		inject.Provide(NewLogger),
		// nats
		inject.Provide(NewNATS),
		// event config
		inject.Provide(func() *natsevent.Config {
			evtConfig := natsevent.Config{
				ClusterID: "test-cluster",
				ClientID:  "nats_client_mons_inject",
				URL:       *natsAddr,
			}
			return &evtConfig
		}),
		// eventer
		inject.Provide(natsevent.NewNATSEventer),
		// domain user service
		inject.Provide(dosvcimpl.NewUserService),
		userServiceBundle,
		// userservice middleware
		inject.Provide(func() []service.UserMiddleware {
			return []service.UserMiddleware{}
		}),
		// todo: service wrapped
		// inject.Provide(service.NewUserService, inject.As(new(service.UserService))),
		// inject.Provide(func(mdw []service.UserMiddleware) service.UserService {
		// 	// todo
		// }),
		// endpoint middleware
		inject.Provide(getEndpointMiddleware),
		// endpoint
		inject.Provide(endpoint.New),
		// server
		inject.Provide(NewServer),
	)
	if err != nil {
		panic(fmt.Sprintf("[Inject] error: %+v", err))
	}
	// di graph
	c.WriteTo(file)

	// logger
	if err := c.Extract(&logger); err != nil {
		panic(fmt.Sprintf("[Logger] error: %+v", err))
	}
	// run
	var db *gorm.DB
	var nc *nats.Conn
	var eventer sharedevent.Eventer
	var eps endpoint.Endpoints
	var g *run.Group
	// nats
	{
		if err := c.Extract(&nc); err != nil {
			panic(fmt.Sprintf("[NATS.Conn] error: %+v", err))
		}
		defer nc.Close()
	}
	// subscribe
	{
		if err := c.Extract(&eventer); err != nil {
			panic(fmt.Sprintf("[Event] error: %+v", err))
		}
		eventSub(eventer)
	}
	// gorm
	{
		if err := c.Extract(&db); err != nil {
			panic(fmt.Sprintf("[gorm.DB] error: %+v", err))
		}
		defer db.Close()
	}
	// logger
	logger.Log("nats.conn", nc.Status())
	// eps
	if err := c.Extract(&eps); err != nil {
		panic(fmt.Sprintf("[Endpoints] error: %+v", err))
	}
	logger.Log("endpoints", fmt.Sprintf("%v", eps))
	if err := c.Extract(&g); err != nil {
		panic(fmt.Sprintf("[Group] error: %+v", err))
	}

	initCancelInterrupt(g)
	log.Println("exit", g.Run())
}
func getEndpointMiddleware(logger kitlog.Logger) (mw map[string][]kitendpoint.Middleware) {
	mw = map[string][]kitendpoint.Middleware{}
	duration := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Help:      "Request duration in seconds.",
		Name:      "request_duration_seconds",
		Namespace: "example",
		Subsystem: "user",
	}, []string{"method", "success"})
	mw["Create"] = []kitendpoint.Middleware{endpoint.LoggingMiddleware(kitlog.With(logger, "method", "Create")), endpoint.InstrumentingMiddleware(duration.With("method", "Create"))}
	// other middleware

	return
}

// server
func NewServer(endpoints endpoint.Endpoints, nc *nats.Conn) (g *run.Group) {
	log.Println("server starting...")
	g = &run.Group{}
	initHTTPHandler(endpoints, g)
	initNATSHandler(endpoints, nc, g)
	// initNATSOverHTTPHandler(endpoints, g)
	initGRPCHandler(endpoints, g)
	return g
}

func initCancelInterrupt(g *run.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(err error) {
		close(cancelInterrupt)
	})
}

func initHTTPHandler(endpoints endpoint.Endpoints, g *run.Group) {
	options := map[string][]kithttp.ServerOption{}
	// Add your HTTP options here

	httpHandler := httphandler.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})
}

/*
func initNATSOverHTTPHandler(endpoints endpoint.Endpoints, g *run.Group) {
	var err error

	nc, err := nats.Connect(*natsAddr)
	if err != nil {
		log.Fatal("nats connection error: ", err)
	}
	httpHandler := natshttp.NewNATSOverHTTPHandler(nc, map[string][]kithttp.ServerOption{})
	httpListener, err := net.Listen("tcp", *natsHTTP)
	if err != nil {
		logger.Log("transport", "NATS over HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "NATS over HTTP", "addr", *natsHTTP)
		return http.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
		nc.Close()
	})
}
*/
func initNATSHandler(endpoints endpoint.Endpoints, nc *nats.Conn, g *run.Group) {
	// log.Println("[initNATSHandler]", "NC", nc, "Group", g)
	var err error
	var findUserSub *nats.Subscription
	var updateUserSub *nats.Subscription

	options := map[string][]kitnats.SubscriberOption{}
	// Add your HTTP options here

	logger.Log("transport", "NATS", "addr", nc.ConnectedUrl())
	natsServer := natshandler.NewNATSServer(endpoints, options)

	findUserSub, err = nc.QueueSubscribe("user.finduser", "user", natsServer.FindUser.ServeMsg(nc))
	if err != nil {
		log.Fatal(err)
	}
	_ = findUserSub
	// defer uSub.Unsubscribe()

	updateUserSub, err = nc.QueueSubscribe("user.updateuser", "user", natsServer.UpdateUser.ServeMsg(nc))
	if err != nil {
		log.Fatal(err)
	}
	_ = updateUserSub
	// defer cSub.Unsubscribe()
}
func initGRPCHandler(endpoints endpoint.Endpoints, g *run.Group) {
	options := map[string][]kitgrpc.ServerOption{}
	// Add your GRPC options here

	grpcServer := grpchandler.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("transport", "GRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "GRPC", "addr", *grpcAddr)
		// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
		// the here demonstrated zipkin tracing middleware.
		// baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
		baseServer := grpc.NewServer()
		pb.RegisterUserServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})
}

// NewDB database
// func NewDB(cfg *iris.Configuration) *gorm.DB {
func NewDB() (*gorm.DB, error) {
	// init db
	db, err := gorm.Open("mysql", *dbConn)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(3 * time.Minute)
	return db, err
}
func NewNATS() (*nats.Conn, error) {
	nc, err := nats.Connect(*natsAddr,
		nats.ReconnectWait(1000*time.Millisecond),
		nats.DisconnectHandler(func(nc *nats.Conn) {
			logger.Log("disconnect", time.Now().Unix())
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logger.Log("reconnect", time.Now().Unix())
		}))
	if err != nil {
		log.Fatal("nats connection error: ", err)
	}
	return nc, err
}
func NewLogger() kitlog.Logger {
	if *debug {
		return kitlog.NewLogfmtLogger(os.Stderr)
	}
	return kitlog.NewNopLogger()
}
func eventSub(eventer sharedevent.Eventer) {
	subFind, err := eventer.Subscribe(event.EventUserFind, "sub.user.find", func(e sharedevent.Event) error {
		log.Printf("[event] subscribe@%v: %+v\n", e.Subject, string(e.Data))
		return nil
	})
	if err != nil {
		log.Printf("[event] sub error: %v", err)
	}
	_ = subFind
	// subUpdating, err := eventer.Subscribe(event.EventUser, "sub.user.updating", func(e sharedevent.Event) error {
	// not support "user.>" express
	subUpdating, err := eventer.Subscribe(event.EventUserUpdating, "sub.user.updating", func(e sharedevent.Event) error {
		log.Printf("[event] subscribe@%v: %+v\n", e.Subject, string(e.Data))
		return nil
	})
	if err != nil {
		log.Printf("[event] sub error: %v", err)
	}
	_ = subUpdating
	subUpdated, err := eventer.Subscribe(event.EventUserUpdated, "sub.user.updated", func(e sharedevent.Event) error {
		log.Printf("[event] subscribe@%v: %+v\n", e.Subject, string(e.Data))
		return nil
	})
	if err != nil {
		log.Printf("[event] sub error: %v", err)
	}
	_ = subUpdated
}
