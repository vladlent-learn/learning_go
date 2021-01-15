package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	x := 123

	// & - "Pointer" to a value
	fmt.Println(&x)

	p := person{
		first: "Sherlock",
		last:  "Holmes",
	}

	fmt.Println(p)

	changeMe(&p)

	fmt.Println(p)

}

func changeMe(p *person) {

	// (*p).first
	p.first = "Professor"
	p.last = "Moriarti"
}
