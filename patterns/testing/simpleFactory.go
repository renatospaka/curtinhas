package testing

import (
	"fmt"

	"github.com/renatospaka/design-patterns/creational/factory"
)

func TestingFactory() {
	fmt.Println("This is the FACTORY implementation")

	ak47, _ := factory.GetGun("ak47")
	musket, _ := factory.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
	fmt.Println("==================================")
	fmt.Println()
}

func printDetails(g factory.GunInterface) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
