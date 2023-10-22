package main

import "fmt"

// thread 1
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	le(hello)
}

// canal apenas RECEBE mensagem (msg<-)
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// canal apenas LÊ mensagem  (<-msg)
func le(data <-chan string) {
	fmt.Println(<- data)
}
