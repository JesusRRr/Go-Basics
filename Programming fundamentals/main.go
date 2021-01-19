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

// Custom type mytype
// underlying type INT
// each type T has and underlying type
type mytype int

func main() {
	fmt.Println("|Programing fundamental|")
	// NUMERIC TYPES

	fmt.Printf("Type: %T\t\t Value %v\n", number0, number0)
	// Declare & assing = INITIALIZATION
	var number1 = 1
	fmt.Printf("Type: %T\t\t Value %v\n", number1, number1)
	// short declaration operator -> ":="
	number2 := 2
	fmt.Printf("Type: %T\t\t Value %v\n", number2, number2)
	// Using Type mytype
	var number3 mytype = 3
	fmt.Printf("Type: %T\t Value %v\n", number3, number3)
	//convertion
	number4 := int(number3) + 1
	fmt.Printf("Type: %T\t\t Value %v\n", number4, number4)
	// uint can be either 32 or 64 bits
	var number5 uint = 5
	fmt.Printf("Type: %T\t\t Value %v\n", number5, number5)
	// int has same size as uint
	var number6 int = 6
	fmt.Printf("Type: %T\t\t Value %v\n", number6, number6)
	// byte is an alias for uint8
	var number7 byte = 7
	fmt.Printf("Type: %T\t\t Value %v\n", number7, number7)
	// rune is an alias for uint32
	// rune is another way to say character
	var number8 rune = 8
	fmt.Printf("Type: %T\t\t Value %v\n", number8, number8)
	// float 32 or 64 bits
	var number9 float32 = 9.000
	fmt.Printf("Type: %T\t\t Value %v\n", number9, number9)
	// complex numbers 64 or 128 bits
	// complex64( float 32 real and imaginary part)
	// complex128( float 64 real and imaginary part)
	var number10 complex64 = 10 + 0i
	fmt.Printf("Type: %T\t\t Value %v\n", number10, number10)

	//STRINGS
	// Strings are inmutable
	// A string is a sequence of bytes
	str1 := "Hello World"
	str1b := []byte(str1)
	fmt.Printf("Type: %T\t\t Value %v\n", str1b, str1b)
	// Strings can be declarated using double quotes
	str2 := "String 2"
	fmt.Printf("Type: %T\t\t Value %v\n", str2, str2)
	// Strings can be declarated using backsticks
	str3 := `"String 3"`
	fmt.Printf("Type: %T\t\t Value %v\n", str3, str3)
	// untype constant
	const str4 = " String 4"
	fmt.Printf("Type: %T\t\t Value %v\n", str4, str4)
	// type constant
	const str5 string = " String 5"
	fmt.Printf("Type: %T\t\t Value %v\n", str5, str5)
}
