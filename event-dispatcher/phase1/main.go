package main

import (
	"github.com/renatospaka/event-dispatcher/phase1/event"
	"github.com/renatospaka/event-dispatcher/phase1/user"
)

func main() {
	ed := event.NewEventDispatcher()
	sendEmailListener := user.NewSendEmailListener()
	publishOnRabbitMQListener := user.NewPublishOnRabbitMQListener()
	ed.AddListener("user_created", sendEmailListener)
	ed.AddListener("user_created", publishOnRabbitMQListener)

	user := user.NewUser(ed)
	_ = user.Create("Renato")
	_ = user.Create("Café")
}
