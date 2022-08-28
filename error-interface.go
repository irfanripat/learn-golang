package main

import (
	"fmt"
	"errors"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	err := errors.New("This is an error")
	if err != nil {
		fmt.Println(err)
	}

	result, err := divide(10, 1)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

