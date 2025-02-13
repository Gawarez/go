package main

import (
	"fmt"
)

func simplesLoop() {
	fmt.Printf("### Simple loop ###\n")
	for i := 0; i <= 100; i += 10 {
		fmt.Println(i)
	}
}

func continueStatement() {
	fmt.Printf("\n### Continue statement ###\n")
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("skipping number:", i)
			continue // goes to next / skip iteration on the loop
		}

		fmt.Println(i)
	}
}

func breakStatement() {
	fmt.Printf("\n### Break statement ###\n")
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("breaking at number:", i)
			break // break / terminate the loop
		}
		fmt.Println(i)
	}
}

func nestedLoops() {
	fmt.Printf("\n### Nested loops ###\n")
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}

	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}

func rangeKeyword() {
	fmt.Printf("\n### Range keyword ###\n")
	// range returns both the index and the value
	fruits := []string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	fmt.Printf("\nOmitting index:\n")
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	fmt.Printf("\nOmitting value:\n")
	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}
}

func main() {
	simplesLoop()
	continueStatement()
	breakStatement()
	nestedLoops()
	rangeKeyword()
}
