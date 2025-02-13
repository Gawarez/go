package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var person_1 Person
	var person_2 Person

	// person_1 specification
	person_1.name = "Hege"
	person_1.age = 45
	person_1.job = "Teacher"
	person_1.salary = 6000

	// person_2 specification
	person_2.name = "Cecilie"
	person_2.age = 24
	person_2.job = "Marketing"
	person_2.salary = 4500

	// Access and print person_1 info
	fmt.Println("### Person 1 ###")
	fmt.Println("Name:", person_1.name)
	fmt.Println("Age:", person_1.age)
	fmt.Println("Job:", person_1.job)
	fmt.Println("Salary:", person_1.salary)

	// Access and print person_2 info
	fmt.Println("\n### Person 2 ###")
	fmt.Println("Name:", person_2.name)
	fmt.Println("Age:", person_2.age)
	fmt.Println("Job:", person_2.job)
	fmt.Println("Salary:", person_2.salary)

	// Pass struct as function arguments
	println("\n### Struct as func args ###")
	printPerson(person_1)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
