package main

import (
	"fmt"
	"time"
)

// thread 1
func main() {
	data := make(chan int)
	workersQtd := 700
	for i:= 0; i < workersQtd; i++ {
		go worker(i, data)
	}

	for i := 0; i < 10000; i++ {
		data <- i
	}
	println()
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
