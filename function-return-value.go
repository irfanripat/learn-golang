package main

import "fmt"

func getHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func getFullName(firstName, lastName string) (string, string) {
	return firstName, lastName
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Irfan"
	middleName = "Ripat"
	lastName = "Ripat"
	return
}

func main() {
	fmt.Println(getHello("Irfan"))
	fmt.Println(getFullName("Irfan", "Ripat"))
	fmt.Println(getCompleteName())
}