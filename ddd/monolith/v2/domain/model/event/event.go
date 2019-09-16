package event

type Event interface {
	Publish(eventName string, args ...interface{})
	Subscribe(eventName string, fn interface{})
}
