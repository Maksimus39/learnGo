package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = append(data, Entry{"Maksim", "Minacov", "89005918911"})
	data = append(data, Entry{"Larisa", "Suslova", "89005918910"})
	data = append(data, Entry{"Andrey", "Minacov", "89205286485"})

	// Различаем команды
	switch arguments[1] {
	// Команда поиска
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	// Команда списка
	case "list":
		list()
		// Ответ на всё остальное
	default:
		fmt.Println("Not a valid option")
	}
}
