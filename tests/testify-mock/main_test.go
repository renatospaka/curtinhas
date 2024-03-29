package main

import (
	"fmt"
	"testing"

	"gopkg.in/stretchr/testify.v1/mock"
)

type smsServiceMock struct {
	mock.Mock
}

// Mocked smsService method
func (m *smsServiceMock) SendChargeNotification(value int) error {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)

	// this records that the method was called and passes in the value
	// it was called with
	args := m.Called(value)

	// it then returns whatever we tell it to return
	// in this case true to simulate an SMS Notification
	// sent out
	return args.Error(0)
}

// we need to satisfy our MessageSerice interface
// which sadly means we have to stub out every method
// defined in that interface
func (m *smsServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

// TestChargeCustomer is where the magic happens
// here we create the SMSService mock
func TestChargeCustomer(t *testing.T) {
	smsService := new(smsServiceMock)

	// then define what should be returned from SendChargeNotification
	// when we pass in the value 100 to it. In this case, we want to return
	// true as it was successful in sending a notification
	smsService.On("SendChargeNotification", 100).Return(nil)

	// next we want to define the service we wish to test
	myService := MyService{smsService}

	// and call said method
	myService.ChargeCustomer(100)

	// at the end, we verify that our mySerice.ChargeCustomer
	// method called our mocked SendChargeNotification method
	smsService.AssertExpectations(t)
}
