package testing

import (
	"fmt"
	abstractFactory "github.com/renatospaka/design-patterns/creational/abstractFactory"
)

func TestingAbstractFactory() {
	fmt.Println("This is the abstract factory implementation")
	adidasFactory, _ := abstractFactory.GetSportsFactory("adidas")
	nikeFactory, _ := abstractFactory.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
	fmt.Println("==================================")
	fmt.Println()
}

func printShoeDetails(s abstractFactory.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s abstractFactory.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
