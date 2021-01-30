package main

import "fmt"

func main() {
	// when we use & before a variable, the value is the address of the variable
	// The type of the address is pointer to the type of the variable (*T)
	// when we use *before an adress we get the value in the address
	a := 43
	b := &a
	*b = 44
	fmt.Printf("Type:%T\t value:%v\n", a, a)
	fmt.Printf("Type:%T\t value:%v\n", b, b)
}
