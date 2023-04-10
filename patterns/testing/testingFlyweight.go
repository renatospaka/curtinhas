package testing

import (
	"fmt"

	"github.com/renatospaka/design-patterns/structural/flyweight"
)

func TestingFlyweight() {
	fmt.Println("This is the FLYWEIGHT implementation")
	game := flyweight.NewGame()

	//Add Terrorist
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(flyweight.CounterTerrroristDressType)
	game.AddCounterTerrorist(flyweight.CounterTerrroristDressType)
	game.AddCounterTerrorist(flyweight.CounterTerrroristDressType)

	dressFactoryInstance := flyweight.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
	fmt.Println("==================================")
	fmt.Println()
}
