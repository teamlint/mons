package handler

import (
	"log"

	"github.com/teamlint/iris"
	"github.com/teamlint/mons/ddd/monolith/v2/application/query"
	"github.com/teamlint/mons/ddd/monolith/v2/application/service"
)

// UserHandler 用户接口层
type UserHandler struct {
	// 应用服务
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (u *UserHandler) GetUser(ctx iris.Context) {
	log.Println("[Adapter.rest] UserHandler.GetUser loading...")
	name := ctx.Params().GetString("name")

	q := query.FindUserQuery{Username: name}
	// concurrency test
	// go u.UserService.FindUser(q)
	// go u.UserService.FindUser(q)
	// go u.UserService.FindUser(q)
	// go u.UserService.FindUser(q)
	// go u.UserService.FindUser(q)
	result, err := u.UserService.FindUser(q)
	if err != nil {
		ctx.WriteString(err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.StopExecution()
		return
	}
	ctx.JSON(result)
}
