package user

import "fmt"

type publishOnRabbitMQListener struct {
	data interface{}
}

func NewPublishOnRabbitMQListener() *publishOnRabbitMQListener {
	return &publishOnRabbitMQListener{}
}

func (l *publishOnRabbitMQListener) SetData(data interface{}) {
	l.data = data
}

func (l *publishOnRabbitMQListener) Handle() error {
	fmt.Println("publicado no RabbitMQ para o usuário:", l.data.(string))
	return nil
}
