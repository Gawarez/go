package main

import (
	"fmt"
)

func familyName(fname string) {
	fmt.Println("Hello,", fname, "Refsnes")
}

func nameAge(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {
	fmt.Println("### Single param ###")
	familyName("Liam")
	familyName("Gabriel")
	familyName("Anja")

	fmt.Println("\n### Multiple params ###")
	nameAge("Liam", 3)
	nameAge("Gabriel", 14)
	nameAge("Anja", 30)
}
