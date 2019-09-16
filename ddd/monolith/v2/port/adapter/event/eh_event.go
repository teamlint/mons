package event

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
	eb "github.com/looplab/eventhorizon/eventbus/local"
	es "github.com/looplab/eventhorizon/eventstore/memory"

	"github.com/teamlint/mons/ddd/monolith/v2/application/service"
	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/event"
)

var _ event.Event = &EhEvent{}

type EhEvent struct {
	Bus        eh.EventBus
	Store      eh.EventStore
	LogService service.LogService
}

func NewEhEvent(ls service.LogService) event.Event {
	// event bus
	bus := eb.NewEventBus(nil)
	go func() {
		for e := range bus.Errors() {
			log.Printf("eventbus: %s", e.Error())
		}
	}()
	// event store
	store := es.NewEventStore()
	e := &EhEvent{Bus: bus, Store: store, LogService: ls}
	e.init()
	return e
}

func (e *EhEvent) Publish(eventName string, args ...interface{}) {
	log.Println("[DomainEvent@eventhorizon] publish " + eventName)
	ctx := context.Background()
	evtType := eh.EventType(eventName)
	edata := make([]interface{}, 0)
	edata = append(edata, args...)
	var evtData eh.EventData = edata
	// evt := eh.NewEvent(evtType, evtData, time.Now())
	aggType := eh.AggregateType("aggtype")
	aggID, _ := uuid.NewUUID()
	version := int(time.Now().Unix())
	evt := eh.NewEventForAggregate(evtType, evtData, time.Now(), aggType, aggID, version)
	if err := e.Bus.PublishEvent(ctx, evt); err != nil {
		log.Printf("[DomainEvent@eventhorizon] publish fail: %v\n", err)
		return
	}
	log.Printf("[event] %# v\n", evt)
	if err := e.Store.Save(ctx, []eh.Event{evt}, version-1); err != nil {
		log.Printf("[error] %v (%T)\n", err, err)
		log.Printf("[DomainEvent@eventhorizon] publish success, but save occur error: %v\n", err)
		return
	}
	log.Println("[DomainEvent@eventhorizon] publish success and store")
}
func (e *EhEvent) Subscribe(eventName string, fn interface{}) {
	log.Println("[DomainEvent@eventhorizon] subscribe " + eventName)
	eventHandler := NewCommonEventHandler(eventName, fn)
	e.Bus.AddObserver(eh.MatchAny(), eventHandler)
}
func (e *EhEvent) init() {
	// e.Subscribe("UserGot", service.LogService.TestEvent)
	e.Subscribe("UserGot", e.LogService.TestEvent)
}

// event handler
type CommonEventHandler struct {
	fn interface{}
}

func NewCommonEventHandler(eventName string, fn interface{}) *CommonEventHandler {
	h := &CommonEventHandler{
		fn: fn,
	}

	return h
}

func (h *CommonEventHandler) HandlerType() eh.EventHandlerType {
	return eh.EventHandlerType("common.eventhandler")
}

func (h *CommonEventHandler) HandleEvent(ctx context.Context, evt eh.Event) error {
	fn := getFunc(h.fn)
	data := evt.Data()
	args := make([]reflect.Value, 0)
	if data != nil {
		evData := data.([]interface{})
		for _, d := range evData {
			args = append(args, reflect.ValueOf(d))
		}
	}
	fn.Call(args)
	return nil
}

// getFunc 获取监听器函数
func getFunc(listener interface{}) reflect.Value {
	fn := reflect.ValueOf(listener)

	if reflect.Func != fn.Kind() {
		panic("不是有效的函数")
	}

	return fn
}
