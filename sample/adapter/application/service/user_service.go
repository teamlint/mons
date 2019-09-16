package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/teamlint/golog"
	"github.com/teamlint/mons/sample/application/command"
	"github.com/teamlint/mons/sample/application/dto"
	"github.com/teamlint/mons/sample/application/event"
	"github.com/teamlint/mons/sample/application/query"
	appsvc "github.com/teamlint/mons/sample/application/service"
	"github.com/teamlint/mons/sample/domain/model"
	"github.com/teamlint/mons/sample/domain/repository"
	"github.com/teamlint/mons/sample/domain/service"
	sharedevent "github.com/teamlint/mons/shared/application/event"
	trans "github.com/teamlint/mons/shared/application/transaction"
	shared "github.com/teamlint/mons/shared/domain/repository"
	"go.uber.org/dig"
)

// UserService user application service
type UserService struct {
	UserRepo    repository.UserRepository
	UserSvc     service.UserService
	RepoContext shared.RepositoryContext
	Eventer     sharedevent.Eventer
}

type UserServiceConfig struct {
	dig.In
	UserRepo    repository.UserRepository
	UserSvc     service.UserService
	RepoContext shared.RepositoryContext
	Eventer     sharedevent.Eventer
}

func NewUserServiceConfig(userRepo repository.UserRepository, userSvc service.UserService, repoContext shared.RepositoryContext, eventer sharedevent.Eventer) UserServiceConfig {
	return UserServiceConfig{
		UserRepo:    userRepo,
		UserSvc:     userSvc,
		RepoContext: repoContext,
		Eventer:     eventer,
	}
}

func NewUserService(conf UserServiceConfig) appsvc.UserService {
	return &UserService{
		UserRepo:    conf.UserRepo,
		UserSvc:     conf.UserSvc,
		RepoContext: conf.RepoContext,
		Eventer:     conf.Eventer,
	}
}
func (s *UserService) Find(_ context.Context, query *query.UserQuery) (*dto.User, error) {
	user, err := s.UserRepo.FindByID(s.RepoContext, query.ID)
	if err != nil {
		return nil, err
	}
	dtoUser := dto.User{}
	if err := dtoUser.From(user); err != nil {
		return nil, err
	}
	// event
	eventData, err := json.Marshal(query)
	if err != nil {
		golog.Errorf("[adapter.UserService.Find] %v", err)
	}
	go s.Eventer.Publish(sharedevent.New(query.ID, event.EventUserType, event.EventUserFind, eventData))
	return &dtoUser, nil

}
func (s *UserService) Update(_ context.Context, cmd *command.UpdateUserCommand) error {
	// event 1
	eventData, err := json.Marshal(cmd)
	if err != nil {
		golog.Errorf("[adapter.UserService.Update] updating  %v", err)
	}
	s.Eventer.Publish(sharedevent.New(cmd.ID, event.EventUserType, event.EventUserUpdating, eventData))
	// transaction
	err = trans.Scope(s.RepoContext, func(rc shared.RepositoryContext) error {
		user := model.User{
			ID:       cmd.ID,
			Username: cmd.Username,
			Nickname: cmd.Nickname,
			Password: cmd.Password,
			Intro:    cmd.Intro,
		}
		return s.UserRepo.Update(rc, &user)
	})
	if err != nil {
		log.Println("[UserService] UpdateUser err:", err)
		return err
	}
	// event 2
	eventUpdatedData, err := json.Marshal(err)
	if err != nil {
		golog.Errorf("[adapter.UserService.Update] updated %v", err)
	}
	s.Eventer.Publish(sharedevent.New(cmd.ID, event.EventUserType, event.EventUserUpdated, eventUpdatedData))
	return nil
}
