package main

import "fmt"

// var keyword
// using var we can declare a variable outside of body
// the scope of this variable is at package level
// if a value is not assigned to a variable, depending the variable type it will be assiged to:
// false for booleans, 0 for integers, 0.0 for floats,"" for strings,
// and nil for pointers, functions, interfaces, slices, channels and maps.
var number0 int

// Declare & assing = INITIALIZATION
var number1 = 1

// Custom type mytype
// underlying type INT
// each type T has and underlying type
type mytype int

var number3 mytype = 3

func main() {
	fmt.Println("Variables, values & type")
	// short declaration operator -> ":="
	number2 := 2
	fmt.Println(number0, number1, number2)
	//formating a string
	fmt.Printf("Type: %T, Value %v\n", number2, number2)
	//formating a string
	fmt.Printf("Type: %T, Value %v\n", number3, number3)
	//convertion
	number4 := int(number3) + 1
	fmt.Printf("Type: %T, Value %v\n", number4, number4)

}
