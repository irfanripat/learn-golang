package main

import "fmt"

func random() interface{} {
	return 1
}


func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)
	
	switch value := result.(type) {
	case string:
		fmt.Println(value, "is String")
	case int:
		fmt.Println(value, "is Int")
	case bool:
		fmt.Println(value, "is Boolean")
	default:
		fmt.Println(value, "is Unknown")
	}	
}