package main

import "fmt"

var metersPerLiter float64

func main() {
	paintNeeDed(4.2, 3.0)
	paintNeeDed(5.9, 5.9)
}

func paintNeeDed(width float64, height float64) {
	metersPerLiter = 10.0
	area := width * height
	fmt.Printf("%.2f  liters needed\n", area/metersPerLiter)
}
