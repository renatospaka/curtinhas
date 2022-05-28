package main

import (
	"github.com/renatospaka/event-dispatcher/event"
	"github.com/renatospaka/event-dispatcher/user"
)

func main() {
	ed := event.NewEventDispatcher()
	sendEmailListener := user.NewSendEmailListener()
	publishOnRabbitMQListener := user.NewPublishOnRabbitMQListener()
	ed.AddListener("user_created", sendEmailListener)
	ed.AddListener("user_created", publishOnRabbitMQListener)

	user := user.NewUser(ed)
	user.Create("Renato")
}
