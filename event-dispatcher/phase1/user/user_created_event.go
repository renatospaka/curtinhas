package user

type UserCreatedEvent struct {
	data string
}

func (e *UserCreatedEvent) GetKey() string {
	return "user_created"
}

func (e *UserCreatedEvent) GetData() interface{} {
	return e.data
}
