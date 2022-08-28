package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Irfan Ripat"
	fmt.Println(strings.Contains(strings.ToLower(name), "ripat"))
}