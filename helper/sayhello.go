package helper

import "fmt"

func SayHello(name string) string{
	fmt.Println("Hello", name, "from package helper")	
	return "Hello " + name
}

func sayGoodBye(name string) {
	fmt.Println("Goodbye", name, "from package helper")
}