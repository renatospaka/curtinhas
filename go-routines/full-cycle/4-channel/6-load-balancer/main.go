package main

import (
	"fmt"
	"time"
)

// thread 1
func main() {
	now := time.Now()

	data := make(chan int)
	workersQtd := 10000
	for i:= 0; i < workersQtd; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1000000; i++ {
		data <- i
	}
	fmt.Println()
	fmt.Printf("com %d workers, demorou %v\n", workersQtd, time.Since(now).String())
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(100*time.Millisecond)
	}
}
