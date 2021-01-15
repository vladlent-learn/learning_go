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

type human interface {
	speak()
}

func main() {
	defer foo()

	sa1 := createAgent("James", "Bond", true)
	sa2 := createAgent("Miss", "Moneypenny", true)

	sa1.speak()

	bar(sa2)
}

func foo() {
	fmt.Println("foo")
}

func bar(h human) {
	h.speak()
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
