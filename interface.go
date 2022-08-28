package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

func sayHello(HasName HasName) {
	fmt.Println("Hello, my name is", HasName.GetName())
}

type Animal struct {
	Name string
}

func (a Animal) GetName() string {
	return a.Name
}

func main() {
	irfan := Person{"Irfan"}
	sayHello(irfan)
	cat := Animal{"Oren"}
	sayHello(cat)
}