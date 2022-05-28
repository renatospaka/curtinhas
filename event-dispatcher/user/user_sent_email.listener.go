package user

import "fmt"

type sendEmailListener struct {
	data interface{}
}

func NewSendEmailListener() *sendEmailListener {
	return &sendEmailListener{}
}

func (l *sendEmailListener) SetData(data interface{}) {
	l.data = data
}

func (l *sendEmailListener) Handle() error {
	fmt.Println("eMail enviado para ", l.data.(string))
	return nil
}
