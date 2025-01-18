package main

import (
	"fmt"
	// "runtime/debug"
	"time"
)

func panicCode() {
	defer handlePanic()

	var x *int
	*x += 1
}

func handlePanic() {
	if panicInfo := recover(); panicInfo != nil {
		fmt.Println("oh no!!", panicInfo)
		// debug.PrintStack()
	} else {
		fmt.Println("never get here")
	}
}

func main() {
	go panicCode()
	time.Sleep(time.Second * 2) // simulates 5 secons of any stuff

	fmt.Println("in this case, it get here...")
}
