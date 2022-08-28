package main

import "fmt"

func getGoodBye(name string) string {
	return fmt.Sprintf("Goodbye %s", name)
}

func main() {
	sayGoodBye := getGoodBye

	fmt.Println(sayGoodBye("Irfan"))
}