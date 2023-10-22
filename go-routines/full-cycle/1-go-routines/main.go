package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i< 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go task("A")
	go task("B")
	go func() {
		for i := 0; i< 14; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(500 * time.Millisecond)
		}
	}()
	time.Sleep(11 * time.Second)
}
