package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teamlint/mons/sample/application/query"
	"github.com/teamlint/mons/sample/application/service"
)

type HomeHandler struct {
	UserSvc service.UserService
}

func NewHomeHandler(userSvc service.UserService) *HomeHandler {
	return &HomeHandler{UserSvc: userSvc}
}
func (h *HomeHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index", gin.H{
		"title": "Home Page",
		"add": func(a int, b int) int {
			return a + b
		},
	})
}
func (h *HomeHandler) About(c *gin.Context) {
	c.HTML(http.StatusOK, "about/index", gin.H{
		"title": "About Page",
	})
}
func (h *HomeHandler) User(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		id = "23f9db62-a145-11e9-ae8f-132379f358df"
	}

	ctx := context.Background()
	query := query.UserQuery{
		ID: id,
	}
	user, err := h.UserSvc.Find(ctx, &query)
	if err != nil {
		c.Error(err)
		return
	}
	c.HTML(http.StatusOK, "home/user.html", user)
}
