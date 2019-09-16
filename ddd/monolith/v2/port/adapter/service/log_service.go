package service

import (
	"log"

	"github.com/teamlint/mons/ddd/monolith/v2/application/service"
)

var _ = service.LogService(&LogService{})

type LogService struct{}

func NewLogService() service.LogService {
	return &LogService{}
}

// func (l *LogService) TestEvent(user *entity.UserInfo) {
func (l *LogService) TestEvent(user interface{}) {
	// func (l *LogService) TestEvent(user string) {
	log.Println("[event] subscribe ", user)
}

// func (l *LogService) HandlerType() eh.EventHandlerType {
// 	return nil
// }

// func (l *LogService) HandleEvent(context.Context, eh.Event) error {
// 	return nil
// }
