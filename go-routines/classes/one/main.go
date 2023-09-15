package main

import (
	"fmt"
	"time"
)

func numeros(done chan <- bool) {
	// chan <- bool = readonly channel
	for i := 0; i<10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}

	done <- true
}

func numeros2(n chan <- int) {
	for i := 0; i<10; i++ {
		n <- i
		fmt.Printf("Escrito no channel: %d\n", i)
		time.Sleep(time.Millisecond * 90)
	}
	fmt.Println("Final da ESCRITA")
	close(n)
}

func letras(done chan <- bool) {
	// chan <- bool = readonly channel
	for l := 'a'; l<'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}

	done <- true
} 

func main() {
	// num := make(chan bool)
	// let := make(chan bool)

	// go numeros(num)
	// go letras(let)

	// <- num
	// <- let

	num2 := make(chan int, 3)
	go numeros2(num2)

	for n := range num2 {
		fmt.Printf("Lido do channel: %d\n", n)
		time.Sleep(time.Millisecond * 180)
	}
	fmt.Println("Final da execução")
}