package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i< 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(200 * time.Millisecond)
		wg.Done()
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(34)

	go task("A", &wg)
	go task("B", &wg)
	go func() {
		for i := 0; i< 14; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}
	}()
	wg.Wait()
}