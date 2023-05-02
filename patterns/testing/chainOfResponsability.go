package testing

import (
	"fmt"

	"github.com/renatospaka/design-patterns/behavioral/chainOfResponsability"
)

func ChainOfResponsability() {
	fmt.Println("This is the chain Of Responsability implementation")
	cashier := &chainOfResponsability.Cashier{}

	//Set next for medical department
	medical := &chainOfResponsability.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &chainOfResponsability.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &chainOfResponsability.Reception{}
	reception.SetNext(doctor)

	patient := &chainOfResponsability.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)

	fmt.Println("==================================")
	fmt.Println()
}
