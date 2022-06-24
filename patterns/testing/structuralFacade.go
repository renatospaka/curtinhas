package testing

import (
	"fmt"
	"log"

	"github.com/renatospaka/design-patterns/structural/facade"
)

func TestingFacade() {
	fmt.Println("This is the facade implementation")
	walletFacade := facade.NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	fmt.Println("==================================")
	fmt.Println()
}
