package main

import "fmt"

func main() {
	const firstName = "Irfan"
	const lastName = "Ripat"

	fmt.Println(firstName, lastName)

	const (
		province = "Makassar"
		country = "Indonesia"
	)

	fmt.Println(province, ",", country)
}