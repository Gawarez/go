package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	cars := [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(cars)

	fmt.Println("Old value:", cars[2])
	cars[2] = "Chevrolet"
	fmt.Println("New value:", cars[2])

	// Print array length
	fmt.Println(len(arr1))
	fmt.Println(len(arr3))
}
