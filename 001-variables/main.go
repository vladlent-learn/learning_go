package main

import "fmt"

// Variable declaration; Has package scope
var x = 42
var y = "James Bond"
var z = true

type mytype int

var myvar mytype

func main() {
	// Shorthand variable declaration; can only be used inside functions
	a := 42
	b := "James Bond"
	c := true

	fmt.Println(a, b, c)
	fmt.Println(x, y, z)

	s := fmt.Sprintf("%d, %s, %t", x, y, z)
	fmt.Println(s)

	fmt.Println("Value of myvar:", myvar)
	fmt.Printf("Type of myvar: %T\n", myvar)
	myvar = 2278
	fmt.Println("Value of myvar:", myvar)

	myvar = mytype(a)
	fmt.Println("Value of myvar:", myvar)
}
