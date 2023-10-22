package main

import "fmt"

// thread 1
func main() {
	ch := make(chan int)
	go publisher(ch)
	reader(ch)
}

func publisher(ch chan int) {
	for i := 0; i<10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}
