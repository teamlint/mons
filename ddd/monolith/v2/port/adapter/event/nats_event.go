package event

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"reflect"

	"github.com/nats-io/nats.go"
	"github.com/teamlint/mons/ddd/monolith/v2/application/service"
	"github.com/teamlint/mons/ddd/monolith/v2/domain/model/event"
)

var _ event.Event = &Event{}

type NatsEvent struct {
	Bus        *nats.Conn
	LogService service.LogService
}

func NewNatsEvent(ls service.LogService) event.Event {
	tokenOpt := nats.Token("abc123")
	nc, err := nats.Connect(nats.DefaultURL, tokenOpt)
	if err != nil {
		log.Fatal("nats connect err:", err)
	}
	// defer nc.Close()
	e := &NatsEvent{Bus: nc, LogService: ls}
	e.init()
	return e
}

func (e *NatsEvent) Publish(eventName string, args ...interface{}) {
	log.Println("[DomainEvent@nats] publish " + eventName)
	msg := nats.Msg{Subject: eventName}
	if len(args) > 0 {
		b, _ := json.Marshal(args[0])
		// b, _ := getBytes(args[0])
		msg.Data = b
	}

	e.Bus.Publish(msg.Subject, msg.Data)
}
func (e *NatsEvent) Subscribe(eventName string, fn interface{}) {
	log.Println("[DomainEvent@nats] subscribe " + eventName)
	hdl := handler(fn)
	e.Bus.Subscribe(eventName, hdl)
}
func (e *NatsEvent) init() {
	// e.Subscribe("UserGot", service.LogService.TestEvent)
	e.Subscribe("UserGot", e.LogService.TestEvent)
}
func handler(fn interface{}) nats.MsgHandler {
	return func(msg *nats.Msg) {
		fn := getFunc(fn)
		data := msg.Data
		args := make([]reflect.Value, 0)
		if data != nil {
			var arg interface{}
			var err error
			err = json.Unmarshal(data, &arg)
			// arg, err = toInterface(data)
			if err != nil {
				log.Printf("[handler] err: %v\n", err)
			}
			args = append(args, reflect.ValueOf(arg))
			// log.Printf("[handler] arg: %v\n", reflect.ValueOf(arg).Interface())
			// args = append(args, reflect.ValueOf(data))
		}
		fn.Call(args)
	}
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func toInterface(buf []byte) (interface{}, error) {
	var data = bytes.NewBuffer(buf)

	var key interface{}
	dec := gob.NewDecoder(data)
	err := dec.Decode(&key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
