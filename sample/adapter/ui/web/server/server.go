package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/teamlint/mons/sample/adapter/application/service"
	"github.com/teamlint/mons/sample/adapter/repository/mysql"
	"github.com/teamlint/mons/sample/adapter/ui/web/handler"
	svc "github.com/teamlint/mons/sample/application/service"
	shared "github.com/teamlint/mons/shared/adapter/repository"

	kitlog "github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	group "github.com/oklog/run"
)

var (
	connStr  = "root:123456@tcp(localhost:3306)/mons?charset=utf8mb4&parseTime=True&loc=Local"
	httpAddr = ":8080"
	logger   kitlog.Logger
)

func Run() {
	// db
	db, err := NewDB()
	if err != nil {
		logger.Log(err)
	}
	defer db.Close()
	// logger
	logger = kitlog.NewLogfmtLogger(os.Stderr)
	// repository
	repoCtx := shared.NewGormRepositoryContext(db)
	repo := mysql.NewUserRepository()
	// service
	svcOpt := service.UserServiceOption{
		UserRepo:    repo,
		RepoContext: repoCtx,
	}
	// svc := service.New([]service.Middleware{})
	svc := service.NewUserService(svcOpt)
	// web server
	g := createServer(svc)

	// interrupt
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())
}

func createServer(userSvc svc.UserService) (g *group.Group) {
	logger.Log("web server starting...")
	g = &group.Group{}
	initHTTPHandler(userSvc, g)
	return g
}

func initCancelInterrupt(g *group.Group) {
	logger.Log("initCancelInterrupt")
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
	}, func(error) {
		close(cancelInterrupt)
	})
}

func initHTTPHandler(userSvc svc.UserService, g *group.Group) {
	homeHandler := handler.NewHomeHandler(userSvc)
	router := gin.Default()
	viewConfig := goview.DefaultConfig
	viewConfig.DisableCache = true
	router.HTMLRender = ginview.New(viewConfig)
	// router.LoadHTMLGlob("views/**/*")
	router.GET("/", homeHandler.Index)
	router.GET("/about", homeHandler.About)
	router.GET("/user/:id", homeHandler.User)
	//
	httpListener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		logger.Log("web server", "HTTP", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("web server", "HTTP", "addr", httpAddr)
		return http.Serve(httpListener, router)
	}, func(error) {
		httpListener.Close()
	})
}

// NewDB database
// func NewDB(cfg *iris.Configuration) *gorm.DB {
func NewDB() (*gorm.DB, error) {
	// init db
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(3 * time.Minute)
	return db, err
}
