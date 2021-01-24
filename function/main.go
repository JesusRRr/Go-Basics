package main

import "fmt"

// function declaration
// func identifier(parameters...) (returns...){body})

// function without parameters and without return
func myFunction1() {
	fmt.Println("myFunction1 was called")
}

//function with parameters but no return
//Everything in go is pass by value
func myFunction2(name string) {
	fmt.Println("myFunction2 was called by", name)
}

//function with parameters and return
func myFunction3(name string) string {
	return fmt.Sprintln("myFunction3 was called by", name)
}

//function with parameters and multiple return
func myFunction4(name string) (string, bool) {
	return fmt.Sprintln("myFunction4 was called by", name), name == "Jesus"
}
func main() {
	// function without arguments
	myFunction1()
	myFunction2("Jesus")
	value := myFunction3("Jesus")
	fmt.Print(value)

	if value, ok := myFunction4("Jesus"); ok {
		fmt.Print(value)
	} else {
		fmt.Println("myFunction4 was not called by", "Jesus")
	}
}
