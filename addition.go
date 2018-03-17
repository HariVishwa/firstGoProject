package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Give first number:")
	fmt.Scan(&x)
	fmt.Print("Give second number:")
	fmt.Scan(&y)

	z := add(x, y)
	fmt.Printf("The addition is:%d", z)
}

// add mperforms addition of two numbers.
func add(x, y int) int {
	return x + y
}

// subtract performs substraction of two numbers.
func subtract(x, y int) int {
	return x - y
}
