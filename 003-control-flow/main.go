package main

import "fmt"

func main() {
	//Same output using different kinds of for loop
	fmt.Println("Alphabet using ASCII:")
	for code := 65; code <= 90; code++ {
		fmt.Printf("%#U \n", code)
	}

	fmt.Println("Alphabet using ASCII:")
	code := 65
	for code <= 90 {
		fmt.Printf("%#U \n", code)
		code++
	}

	fmt.Println("Alphabet using ASCII:")
	code = 65
	for {
		fmt.Printf("%#U \n", code)
		code++
		if code == 90 {
			break
		}
	}
}

// Switch examples
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
