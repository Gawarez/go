package main

import (
	"fmt"
)

func definingMaps() {
	// Ways of creating a map
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	fmt.Printf("a\t%v\n", a)

	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	fmt.Printf("\nb\t%v\n", b)

	c := make(map[string]int)
	c["Oslo"] = 1
	c["Bergen"] = 2
	c["Trondheim"] = 3
	c["Stavanger"] = 4
	fmt.Printf("\nc\t%v\n\n", c)

	// empty map
	var d = make(map[string]string)
	fmt.Println(d == nil)
	var e map[string]string
	fmt.Println(e == nil)
}

func mapElements() {
	// Access map elements
	var car = make(map[string]string)
	car["brand"] = "Ford"
	car["model"] = "Mustang"
	car["year"] = "1964"
	fmt.Println("Access map elements:", car)

	// Update and add map elements
	car["year"] = "1970"
	car["color"] = "red"

	fmt.Println("\nUpdate and add element:", car)

	// Delete element from map
	delete(car, "year")
	fmt.Println("\nDelete element:", car)
}

func checkElements() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := a["brand"] // Checking for existing key and its value
	val2, ok2 := a["color"] // Checking for non-existing key and its value
	val3, ok3 := a["day"]   // Checking for existing key and its value
	_, ok4 := a["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
}

func referenceMap() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1970"
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)
}

func iterateOverMap() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b []string // defining the order
	b = append(b, "one", "two", "three", "four")

	for k, v := range a { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b { // loop with the defined order
		fmt.Printf("%v : %v, ", element, a[element])
	}
}

func main() {
	// definingMaps()
	// mapElements()
	// checkElements()
	// referenceMap()
	iterateOverMap()
}
