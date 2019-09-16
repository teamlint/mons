package rest

import (
	"log"

	"github.com/teamlint/iris"
	"github.com/teamlint/mons/ddd/monolith/facade/application/service"
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
	log.Println("[Facade] UserHandler.GetUser loading...")
	name := ctx.Params().GetString("name")
	result, err := u.UserService.GetUserByName(name)
	if err != nil {
		ctx.WriteString(err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.StopExecution()
		return
	}
	ctx.JSON(result)
}
