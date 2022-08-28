package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khanedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		1,3,
	}

	fmt.Println(values)
}