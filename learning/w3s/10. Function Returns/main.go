package main

import (
	"fmt"
)

// Receives 2 integers and returns as integer (force case)
func myFunction(x int, y int) int {
	return x + y
}

// Named return values
func namedReturn(x int, y int) (result int) {
	result = x + y
	return
}

// Multiple return values
func multReturn(x int, y string) (result int, txt string) {
	result = x + x
	txt = y + " World!"
	return
}

func main() {
	fmt.Println(myFunction(1, 2))
	fmt.Println()
	fmt.Println(namedReturn(1, 2))
	fmt.Println()
	fmt.Println(multReturn(5, "Hello"))
}
