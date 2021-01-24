package main

import "fmt"

var name string

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

//variadic parmaters
//The output is a slice of the parameter type
//The variadic parmather must be the end paramater
func setLanguajes(name string, languajes ...string) []string {
	fmt.Println("Function setLanguajes was called by", name)
	return languajes
}

func main() {
	name = "Jesus"
	// function without arguments
	myFunction1()
	myFunction2(name)
	value := myFunction3(name)
	fmt.Print(value)

	if value, ok := myFunction4(name); ok {
		fmt.Print(value)
	} else {
		fmt.Println("myFunction4 was not called by", name)
	}
	languajes := setLanguajes(name, "go", "java", "python")
	fmt.Printf("Type: %T\t Values: %v\n", languajes, languajes)
	fmt.Println("Languajes' lenght", len(languajes))
	fmt.Println("Languajes' capacity", cap(languajes))

	//we can pass a slice as a variadic argument
	anotherLan := []string{"kotlin", "javascript", "c++"}
	languajes = append(languajes, setLanguajes(name, anotherLan...)...)
	fmt.Printf("Type: %T\t Values: %v\n", languajes, languajes)
	fmt.Println("Languajes' lenght", len(languajes))
	fmt.Println("Languajes' capacity", cap(languajes))
}
