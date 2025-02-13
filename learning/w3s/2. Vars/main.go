package main

import (
	"fmt"
)

func main() {
	// var a string = "John"
	// var b int = 1
	// var c bool = true
	// d := "strings"

	// var a, b = 6, "Hello"
	// c, d := 7, "World!"

	// same line
	// var a, b, c, d int = 1, 3, 5, 7

	// Block
	var (
		a int
		b int = 1
		c string = "hello"
		d = 21
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
