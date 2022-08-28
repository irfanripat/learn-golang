package main

import "fmt"

func main() {

	var person map[string]string = map[string]string{
		"name": "Irfan",
		"address": "Makassar",
	}

	fmt.Println(person)

	person["age"] = "18"
	fmt.Println(person["age"])

	var book = make(map[string]string)

	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"

	fmt.Println(book)
} 