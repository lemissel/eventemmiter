package eventemmiter

type Eventlistener interface {
	Subscribe(eventName string, callback func(args ...interface{}))
	Emit(eventNameToDispatch string, args ...interface{})
	Remove(eventNameToRemove string)
}

type eventListenerImpl struct {
	events events
}

type events map[string]func(args ...interface{})

func New() *eventListenerImpl {
	return &eventListenerImpl{
		events: events{},
	}
}

// Subscribe adds a new event to events list
func (el *eventListenerImpl) Subscribe(eventName string, callback func(args ...interface{})) {
	el.events[eventName] = callback
}

// Emit execute event found
func (el *eventListenerImpl) Emit(eventNameToDispatch string, args ...interface{}) {
	if event, found := el.events[eventNameToDispatch]; found {
		event(args)
	}
}

// Remove removes an event in the events list
func (el *eventListenerImpl) Remove(eventNameToRemove string) {
	delete(el.events, eventNameToRemove)
}
