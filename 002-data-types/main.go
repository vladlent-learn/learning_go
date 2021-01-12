package main

import "fmt"

const (
	// Typed constant
	LEET int = 1337

	// Untyped constant
	PI = 3.14
)

const (
	NEXT_YEAR  = iota + 2022
	NEXT_YEAR1 = iota + 2022
	NEXT_YEAR2 = iota + 2022
	NEXT_YEAR3 = iota + 2022
)

func main() {
	n := 228

	// Shifted one bit to the left
	x := n << 1

	str := `	Created 
	a string 
	using string 
	literal`

	fmt.Println("decimal - binary - hex")
	fmt.Printf("%d - %b -  %#x\n", n, n, n)
	fmt.Printf("%d - %b -  %#x\n", x, x, x)

	fmt.Println(str)
	fmt.Println("Constants:", LEET, PI, NEXT_YEAR, NEXT_YEAR1, NEXT_YEAR2, NEXT_YEAR3)
}
