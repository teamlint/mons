package eventstore

import "log"

type Event struct {
}

func (e *Event) Publish(eventName string) {
	log.Println("[ApplicationService] event Publish " + eventName)
}
func (e *Event) Subscribe(eventName string) {
	log.Println("[ApplicationService] event subscribe " + eventName)
}
