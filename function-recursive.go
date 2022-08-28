package main

import "fmt"

func factorialLoop(value int) int {
	var result = 1
	for i := 1; i <= value; i++ {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorialRecursive(value - 1)
}

func main() {
	loop := factorialLoop(5)
	recursive := factorialRecursive(5)
	fmt.Println(loop)
	fmt.Println(recursive)
	fmt.Println(5*4*3*2*1)
}