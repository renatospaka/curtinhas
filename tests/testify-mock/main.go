package main

import "fmt"

// MessageService handles notifying clients they have
// been charged
type MessageService interface {
	SendChargeNotification(int) error
}

// SMSService is our implementation of MessageService
type SMSService struct {}


// MyService uses the MessageService to notify clients
type MyService struct {
	messageService MessageService
}

func (s SMSService) SendChargeNotification(value int) error {
	fmt.Println("Sending Production Charge Notification")
	return nil
}

func (a MyService) ChargeCustomer(value int) error {
	a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging Customer For the value of %d\n", value)
	return nil
}

func main() {
	fmt.Println("Hello World")

	smsService := SMSService{}
	myService := MyService{smsService}
	myService.ChargeCustomer(100)
}
