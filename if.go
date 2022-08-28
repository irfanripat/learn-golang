package main

import "fmt"

func main() {
	var name = "Irfan"

	if name == "Ifan" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Who are you?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Mantep")
	}
}