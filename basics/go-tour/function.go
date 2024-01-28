package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

// return multiple values
func swap(x, y string) (string, string) {
	return y, x
}

// returning value with empty return statement
func nakedReturn() (c int) {
	c = 10
	return
}

func main() {
	fmt.Print("Hello World\n")
	fmt.Printf("Value of PI is %g\n", math.Pi)
	fmt.Print(add(2, 3), "\n")
	a, b := swap("Wick", "John")
	fmt.Print(a, " ", b, "\n")

	fmt.Print(nakedReturn(), "\n")
}
