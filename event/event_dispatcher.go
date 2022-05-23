package event

type Listener interface {
	SetData(data interface{})
	Handle() error
}

type Event interface {
	GetKey() string
	GetData() interface{}
}

type eventDispatcher struct {
	Listeners map[string][]Listener
}

func NewEventDispatcher() *eventDispatcher {
	return &EventDispatcher{
		Listeners: make(map[string][]Listener),
	}
}

func (e *EeventDispatcher) AddListener(event strng, listener Listener) {
	if e.Listeners == nil {
		e.Listeners = make(map[string][]Listener)
	}
	e.Listeners[event] = append(e.Listeners[event], listener)
}

func (e *eventDispatcher) Dispatch(event Event) {
	if e.Listeners == nil {
		return
	}

	for _, listener := range e.Listeners[event.GetKey()] {
		listener.SetData(event.GetData())
		listener.Handle()
	}
}