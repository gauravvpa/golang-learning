package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello World\n")
	// variable declarations

	// 1. var with data type
	var isActive bool
	var i int
	fmt.Print(isActive, i, "\n")

	// 2. When Initialize is present can skip the data type
	var isToday = true
	fmt.Print(isToday, "\n")

	// 3. Short Variable declaration Only Inside the function
	testVar := 5
	fmt.Print(testVar, "\n")

	// 4. multiple variable declaration
	var (
		name1 int
		name2 string = "Test"
	)
	fmt.Print(name1, name2, "\n")

	// Type of varaible
	fmt.Printf("%T", testVar)

	// when we don't specify variable type, the type is derived from the value on right hand side
	v := 0
	fmt.Printf("v is of Type %T\n", v)

	// constants
	const constVar = 4 + 9i
}
