package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	defer foo()

	sa1 := createAgent("James", "Bond", true)
	sa2 := createAgent("Miss", "Moneypenny", true)

	sa1.speak()
	sa2.speak()

	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func createAgent(first string, last string, ltk bool) secretAgent {
	return secretAgent{
		person{
			first,
			last,
		},
		ltk,
	}
}

func (s secretAgent) speak() {
	fmt.Println("My name is", s.last+",", s.first, s.last)
}
