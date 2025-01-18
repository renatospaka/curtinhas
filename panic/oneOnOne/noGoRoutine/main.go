package main

import "fmt"

func panicCode() {
	var x *int
	*x += 1
}

func handlePanic() {
	if panicInfo := recover(); panicInfo != nil {
		fmt.Println("oh no!!", panicInfo)
	} else {
		fmt.Println("never goes here")
	}
}

func main() {
	defer handlePanic()

	panicCode()
	fmt.Println("never get here either...")
}
