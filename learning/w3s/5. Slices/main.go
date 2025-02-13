package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Slice exampple 1:\n")
	sliceExample1 := []int{}
	fmt.Println(len(sliceExample1))
	fmt.Println(cap(sliceExample1))
	fmt.Println(sliceExample1)

	sliceExample2 := []string{"Go", "slices", "are", "powerful!"}
	fmt.Printf("\nSlice example 2:\n")
	fmt.Println(len(sliceExample2))
	fmt.Println(cap(sliceExample2))
	fmt.Println(sliceExample2)

	// Slice from an array
	fmt.Printf("\nSlice from an Array:\n")
	sliceFromArray := sliceExample2[1:4]
	fmt.Println(sliceFromArray)
	fmt.Printf("Capacity = %d\n", cap(sliceFromArray))

	// Create a slice with the make() function
	fmt.Printf("\nSlice from make() func:\n")
	sliceMake1 := make([]int, 5, 10)
	fmt.Printf("sliceMake1 = %v\n", sliceMake1)
	fmt.Printf("length = %v\n", len(sliceMake1))
	fmt.Printf("capacity = %v\n", cap(sliceMake1))

	// omitted capacity
	sliceMake2 := make([]int, 5)
	fmt.Printf("\nsliceMake2 = %v\n", sliceMake2)
	fmt.Printf("length = %v\n", len(sliceMake2))
	fmt.Printf("capacity = %v\n", cap(sliceMake2))
}
