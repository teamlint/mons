package service

type LogService interface {
	// TestEvent(user *entity.UserInfo)
	TestEvent(user interface{})

	// // HandlerType is the type of the handler.
	// HandlerType() eh.EventHandlerType

	// // HandleEvent handles an event.
	// HandleEvent(context.Context, eh.Event) error
}
