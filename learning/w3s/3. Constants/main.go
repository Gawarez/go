package main

import (
	"fmt"
)

// Constant names written in uppercase
const PI = 3.14

// Block
const (
  A int8 = 1
  B = 3.14
  C = "Hi!"
)

func main() {
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
