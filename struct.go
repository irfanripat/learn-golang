package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, my name is", customer.Name)
}

func main() {
	c := Customer{"Joe", "123 Main", 32}
	c.Name = "John"
	c.Address = "124 Main"
	c.Age = 33
	fmt.Println(c.Name)
	c.sayHello()
}