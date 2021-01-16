package main

import "fmt"

func main() {
	// Both sender and receiver
	c := make(chan int)

	// Send-only
	sender := make(chan<- int)

	// Receiver-only
	receiver := make(<-chan int)

	go func() {
		c <- 42
		c <- 45
	}()

	go send(c)
	receive(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("---------------")
	fmt.Printf("Bothway %T\n", c)
	fmt.Printf("Sender %T\n", sender)
	fmt.Printf("Receiver %T\n", receiver)
}

func send(c chan<- int) {
	c <- 228
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
