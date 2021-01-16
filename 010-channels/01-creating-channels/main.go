package main

import "fmt"

func main() {
	c := make(chan int)

	// make(chan int, 2) - will create a buffered channel

	go func() {
		c <- 42
		c <- 45
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
}
