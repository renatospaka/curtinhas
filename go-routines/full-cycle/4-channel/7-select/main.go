package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	msg string
}

// thread 1
func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			// time.Sleep(time.Second * 1)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			// time.Sleep(time.Second * 2)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	// for i := 0; i < 2; i++ {
	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("received from RabbitMQ: %d\n", msg1.id)

		case msg2 := <-c2:
			fmt.Printf("received from Kafka: %d\n", msg2.id)

		case <-time.After(time.Second * 3):
			println("timeout")

		// default:
		// 	// executa primeiro pq as outras 3 condições estão "em espera"
		// 	println("default")
		}
	}
}
