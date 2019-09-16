package service

import (
	"log"

	"github.com/google/uuid"
	"github.com/teamlint/mons/ddd/monolith/v2/application/query"
	"github.com/teamlint/mons/ddd/monolith/v2/application/transaction"
	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/account/entity"
	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/account/repository"
	doservice "github.com/teamlint/mons/ddd/monolith/v2/domain/model/account/service"
	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/event"
	shared "github.com/teamlint/mons/ddd/monolith/v2/domain/model/shared/repository"
	"go.uber.org/fx"
)

// UserService 用户应用服务
type UserService struct {
	DomainUserService *doservice.UserService
	UserRepository    repository.UserRepository
	Context           shared.RepositoryContext
	Event             event.Event
}
type UserServiceParams struct {
	fx.In
	DomainUserService *doservice.UserService
	UserRepository    repository.UserRepository
	Context           shared.RepositoryContext
	Event             event.Event
	// Event             event.Event `name:"event"`
}

// func NewUserService(svc *doservice.UserService, ctx shared.RepositoryContext, repo repository.UserRepository, e event.Event) *UserService {
// return &UserService{DomainUserService: svc, Context: ctx, UserRepository: repo, Event: e}
// }
func NewUserService(params UserServiceParams) *UserService {
	return &UserService{
		DomainUserService: params.DomainUserService,
		Context:           params.Context,
		UserRepository:    params.UserRepository,
		Event:             params.Event,
	}
}

func (u *UserService) FindUser(q query.FindUserQuery) (*entity.UserInfo, error) {
	log.Println("[ApplicationService] UserService.GetUserByName loading")
	var err error
	user, err := u.DomainUserService.GetUserByName(u.Context, q.Username)
	if err != nil {
		// u.Context.Rollback()
		return nil, err
	}
	roles, err := u.UserRepository.GetUserRoles(u.Context, user.ID)
	if err != nil {
		// u.Context.Rollback()
		return nil, err
	}
	// create user
	id, _ := uuid.NewUUID()
	newUser := entity.User{
		ID:       id.String(),
		Username: "abc",
	}
	// transaction scope
	err = transaction.Scope(u.Context, func(rc shared.RepositoryContext) error {
		log.Println("[ApplicationService] UserService.GetUserByName transaction scope")
		err = u.UserRepository.CreateUser(rc, &newUser)
		if err != nil {
			return err
			// return nil, err
		}
		id2, _ := uuid.NewUUID()
		newUser.ID = id2.String()
		newUser.Username = "user 2"
		err = u.DomainUserService.CreateUser(rc, &newUser)
		// err = fmt.Errorf("[test rollback]")
		log.Println(err)
		// return nil, err
		return err
	})
	// get user
	if err != nil {
		return nil, err
	}
	userInfo := entity.UserInfo{User: *user, Roles: roles}
	// event
	u.Event.Publish("UserGot", &userInfo)

	return &userInfo, nil
}
