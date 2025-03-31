package main

import (
	"fmt"
	"strings"
)

func main() {

	// Вызов даты года

	//var now time.Time = time.Now()
	//var year int = now.Year()
	//
	//fmt.Println(year)

	// Вызов методов продолжение

	broken := "G# r#cks"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
