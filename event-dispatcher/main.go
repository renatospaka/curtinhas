package main

import (
	"github.com/renatospaka/event-dispatcher/event"
	"github.com/renatospaka/event-dispatcher/user"
)

func main() {
	ed := event.NewEventDispatcher()
	sendEmailListener := user.NewSendEmailListener()
	ed.AddListener("user_created", sendEmailListener)
	
	user := user.NewUser(ed)
	user.Create("Renato")
}
