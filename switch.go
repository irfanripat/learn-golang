package main

import "fmt"

func main() {
	var name = "Irfan"

	switch name {
	case "Irfan":
		fmt.Println("Hai Irfan")
	default:
		fmt.Println("Kenalan dong")
	}

	switch len(name) {
	case 5:
		fmt.Println("Mantep")
	default:
		fmt.Println("Kepanjangan")
	}
}