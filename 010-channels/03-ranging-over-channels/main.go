package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)
	receive(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
