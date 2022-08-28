package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name)
}

func main() {
	irfan := Man{"Irfan"}
	irfan.Married()

	fmt.Println(irfan.Name)
}