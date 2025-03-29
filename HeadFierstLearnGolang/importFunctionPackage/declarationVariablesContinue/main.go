package main

import (
	"fmt"
	"reflect"
)

// Объявление переменной и присвание значения
var quantity = 4
var length, width = 1.2, 2.4
var customerName = "Gopher"

// Нулевое значение переменной
var myInt int
var myFloat float64
var muBool bool

func main() {
	// type
	fmt.Println(reflect.TypeOf(quantity))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(customerName)

	// type
	fmt.Println(myInt)
	fmt.Println(myFloat)
	fmt.Println(muBool)
}
