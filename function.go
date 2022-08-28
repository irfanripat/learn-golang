package main

import "fmt"

func main() {
	sayHello("Irfan")
}

func sayHello(name string) {
	fmt.Println("Hello", name)
}