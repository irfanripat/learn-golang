package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args
	fmt.Println(args)

	name, error := os.Hostname()
	if error == nil {
		fmt.Println(name)
	} else {
		fmt.Println("Error:", error)
	}

	username := os.Getenv("APP_USERNAME")
	fmt.Println(username)

}