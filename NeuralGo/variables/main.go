package main

import "fmt"

var middleName = "Key"

func main() {
	age := 10
	middleName = "Value"
	fmt.Println(middleName)

	fmt.Println(age)

	count := 20

	fmt.Println(count)

	lastname := "Minakov"
	fmt.Println(lastname)

	// Default values
	// Pointers, slices, maps, function, and struct

	// ---- SCOPES
	fmt.Println(middleName)

	fmt.Println(lastNameLarisa())
}

func lastNameLarisa() string {
	name := "Larisa"
	return name
}
