package event

import (
	"log"

	eb "github.com/asaskevich/EventBus"
	"github.com/teamlint/mons/ddd/monolith/v2/application/service"
	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/event"
)

var _ event.Event = &Event{}

type Event struct {
	Bus        eb.Bus
	LogService service.LogService
}

func New(ls service.LogService) event.Event {
	bus := eb.New()
	e := &Event{Bus: bus, LogService: ls}
	e.init()
	return e
}

func (e *Event) Publish(eventName string, args ...interface{}) {
	log.Println("[DomainEvent] publish " + eventName)
	e.Bus.Publish(eventName, args...)
}
func (e *Event) Subscribe(eventName string, fn interface{}) {
	log.Println("[DomainEvent] subscribe " + eventName)
	e.Bus.Subscribe(eventName, fn)
}
func (e *Event) init() {
	// e.Subscribe("UserGot", service.LogService.TestEvent)
	e.Subscribe("UserGot", e.LogService.TestEvent)
}
