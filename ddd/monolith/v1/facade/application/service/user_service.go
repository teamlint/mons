package service

import (
	"log"

	"github.com/teamlint/mons/ddd/monolith/facade/application/eventstore"
	"github.com/teamlint/mons/ddd/monolith/facade/application/service/domain/model"
	doservice "github.com/teamlint/mons/ddd/monolith/facade/application/service/domain/service"
)

// UserService 用户应用服务
type UserService struct {
	DomainUserService *doservice.UserService
	event             *eventstore.Event
}

func NewUserService(svc *doservice.UserService) *UserService {
	return &UserService{DomainUserService: svc, event: new(eventstore.Event)}
}

func (u *UserService) GetUserByName(name string) (*model.User, error) {
	log.Println("[ApplicationService] UserService.GetUserByName loading...")
	u.event.Publish("GetUserByName:" + name)
	return u.DomainUserService.GetUserByName(name)
}
