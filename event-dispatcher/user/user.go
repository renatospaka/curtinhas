package user

import "github.com/renatospaka/event-dispatcher/event"

type user struct {
	name string
	dispatcher *event.EventDispatcher
}

func NewUser(dispatcher *event.EventDispatcher) *user {
	return &user{
		dispatcher: dispatcher,
	}
}

func (u *user) Create(name string) error {
	u.name = name
	userCreatedEvent := &UserCreatedEvent{data: name}
	u.dispatcher.Dispatch(userCreatedEvent)

	return nil
}
