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

func letras(done chan <- bool) {
	// chan <- bool = readonly channel
	for l := 'a'; l<'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}

	done <- true
} 

func main() {
	num := make(chan bool)
	let := make(chan bool)

	go numeros(num)
	go letras(let)

	<- num
	<- let

	fmt.Println("Final da execução")
}