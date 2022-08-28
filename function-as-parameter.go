package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println(filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "Asshole"
	} else {
		return "Hello " + name
	}
}

func upperCharFilter(name string) string {
	return strings.ToUpper(name)
}

func main() {
	sayHelloWithFilter("Irfan", spamFilter)
	sayHelloWithFilter("Anjing", upperCharFilter)
}