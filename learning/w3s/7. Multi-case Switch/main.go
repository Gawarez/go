package main

import (
	"fmt"
)

func main() {
	day := 1

	// Multi-case switch
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}
}
