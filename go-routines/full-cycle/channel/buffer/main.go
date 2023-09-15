package main

// thread 1
func main() {
	// // causa um deadlock
	// ch1 := make(chan string)
	// ch1 <- "hello"
	// ch1 <- "hello"
	// println(<-ch1)
	// println(<-ch1)
	
	// NÃO causa um deadlock
	ch2 := make(chan string, 2)
	ch2 <- "hello"
	ch2 <- "hello"
	println(<-ch2)
	println(<-ch2)
}
