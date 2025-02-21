package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var result []int
	seen := make(map[int]bool)

	for len(result) < 6 {
		num, _ := rand.Int(rand.Reader, big.NewInt(61))
		val := int(num.Int64())

		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}

	fmt.Println(result)
}
