package main

import (
	"fmt"
	"sync"
)

// thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publisher(ch)
	go reader(ch, &wg)

	wg.Wait()
}

func publisher(ch chan int) {
	for i := 0; i<10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}
