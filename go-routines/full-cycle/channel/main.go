package main

import (
	"fmt"
	"time"
)

// thread 1
func main() {
	canal := make(chan string)

	// thread 2
	go func() {
		time.Sleep(500 * time.Millisecond)
		canal <- "Olá canal" //enche o canal
	}()

	// thread 1
	msg := <- canal	// esvazia o canal
	fmt.Println(msg)
}
