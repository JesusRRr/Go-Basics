package main

import (
	"fmt"
)

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

// uint can be either 32 or 64 bits
var number5 uint = 5

// int has same size as uint
var number6 int = 6

// byte is an alias for uint8
var number7 byte = 7

// rune is an alias for uint32
// rune is another way to say character
var number8 rune = 8

// float 32 or 64 bits
var number9 float32 = 9.000

// complex numbers 64 or 128 bits
// complex64( float 32 real and imaginary part)
// complex128( float 64 real and imaginary part)
var number10 complex64 = 10 + 0i

func main() {
	fmt.Println("|Variables, values & type|")
	// short declaration operator -> ":="
	number2 := 2
	fmt.Printf("Type: %T, Value %v\n", number0, number0)
	fmt.Printf("Type: %T, Value %v\n", number1, number1)
	fmt.Printf("Type: %T, Value %v\n", number2, number2)
	fmt.Printf("Type: %T, Value %v\n", number3, number3)
	//convertion
	number4 := int(number3) + 1
	fmt.Printf("Type: %T, Value %v\n", number4, number4)
	fmt.Printf("Type: %T, Value %v\n", number5, number5)
	fmt.Printf("Type: %T, Value %v\n", number6, number6)
	fmt.Printf("Type: %T, Value %v\n", number7, number7)
	fmt.Printf("Type: %T, Value %v\n", number8, number8)
	fmt.Printf("Type: %T, Value %v\n", number9, number9)
	fmt.Printf("Type: %T, Value %v\n", number10, number10)

}
