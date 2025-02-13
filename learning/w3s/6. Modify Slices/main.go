package main

import "fmt"

func main() {
	prices := []int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	a := [3]int{1, 5, 20}
	b := [...]int{1, 5, 20}
	fmt.Println(a == b)
}
