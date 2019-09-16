package main

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/teamlint/iris"
	"github.com/teamlint/mons/ddd/monolith/facade/application/service"
	"github.com/teamlint/mons/ddd/monolith/facade/application/service/impl"
	"github.com/teamlint/mons/ddd/monolith/facade/application/service/repository"
	"github.com/teamlint/mons/ddd/monolith/facade/application/service/repository/mysql"
	"github.com/teamlint/mons/ddd/monolith/facade/rest"
	"github.com/teamlint/mons/ddd/monolith/presentation/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			NewDB,
			NewUserRepository,
			NewUserService,
			NewUserHandler,
			NewConfig,
			server.New,
		),
		fx.Invoke(server.Bootstrap),
	)
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

// NewUserRepository user repository
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	userRepo := mysql.NewUserRepository(db)
	return userRepo
}

// NewUserService user application service
func NewUserService(repo repository.UserRepository) service.UserService {
	// init service
	userService := impl.NewUserService(repo)
	return userService
}
func NewUserHandler(userService service.UserService) *rest.UserHandler {
	// init handler
	userHandler := rest.NewUserHandler(userService)
	return userHandler
}
func NewConfig() *iris.Configuration {
	cfg := iris.YAML("./config.yml")
	return &cfg
}
