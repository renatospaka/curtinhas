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
	fmt.Println("publicado no RabbitMQ. Nome do usuário: ", l.data.(string))
	return nil
}
