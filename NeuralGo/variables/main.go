package main

import "fmt"

var midddleName = "Key"

func main() {
	age := 10

	fmt.Println(age)

	count := 20

	fmt.Println(count)

	lastname := "Minakov"
	fmt.Println(lastname)

	// Default values
	// Pointers, slices, maps, function, and struct

	// ---- SCOPES
	fmt.Println(midddleName)

	fmt.Println(lastNameLarisa())
}

func lastNameLarisa() string {
	name := "Larisa"
	return name
}
