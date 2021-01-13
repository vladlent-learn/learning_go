package main

import "fmt"

func main() {
	arr := [5]int{6, 223, 3456, 45, 90}

	// Iterating over array using range which yields index and value on each iteration
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("The type of array is: %T\n", arr)

	// Creating a slice using slice literal
	x := []int{4, 5, 8, 12}
	fmt.Println(len(x))

	z := make([]int, 10, 100)
	fmt.Println(z)

	// Iterating over slice using range which yields index and value on each iteration
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("The type of slice is: %T\n", z)

	// Slicing a slice
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:3])

	// Append to a slice
	y := append(x, 1, 2, 3)
	fmt.Println(y)

	// Append slice to a slice
	y = append(y, y...)
	fmt.Println(y)

	// Deleting from a slice at position 3
	y = append(y[:2], y[4:]...)
	fmt.Println(y)

	// Creating a map with map literal
	// Part between = and { is the TYPE
	m := map[string]int{
		"foo": 32,
		"bar": 1234,
	}

	fmt.Println(m)

	fmt.Println(m["foo"])

	v, ok := m["baz"]
	fmt.Println(v)
	fmt.Println(ok)

	// Checking that value with corresponding key exists using "comma ok" idiom
	if v, ok := m["abc"]; ok {
		fmt.Println(v)
	}

	if v, ok := m["bar"]; ok {
		fmt.Println(v)
	}

	// Adding a value to the map
	m["hello"] = 487
	fmt.Println(m)

	// Iterating over a map with range which yields key and value on each iteration
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Deleting from map
	delete(m, "hello")
	fmt.Println(m)
}
