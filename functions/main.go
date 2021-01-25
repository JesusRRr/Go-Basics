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

//struct developer
type developer struct {
	name    string
	languje string
}

//struct dataScientist
type dataScientist struct {
	name    string
	languje string
}

//Methods
//func ( r receiver) identifier(parameters) (returns){body}
func (dev developer) develop() {
	fmt.Println(dev.name, "is developing in", dev.languje, "languaje")
}

func (ds dataScientist) develop() {
	fmt.Println(ds.name, "uses", ds.languje, "for data Science")
}

//Interfaces
//Any type implements the method in a interface, it is now type interface
//Every type implements an empty interface: interface{}
type person interface {
	develop()
}

//This method uses the person interface as a parameter
//Any type that implements develop method, could be passed as an argument
//This is polymorphism
func speak(per person) {
	var name string
	var role string
	switch per.(type) {
	case developer:
		name = per.(developer).name
		role = "Developer"
	case dataScientist:
		name = per.(dataScientist).name
		role = "Data Scientist"
	}
	fmt.Printf("I am %v and I am %v\n", name, role)
}

// return a function
func x() func() string {
	// Anomymous function
	return func() string {
		return "This code comes from x"
	}
}

func main() {
	name = "Jesus"
	// function without arguments
	// defer keyword will run after main func runs
	defer myFunction1()
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

	jesusDev := developer{
		name:    name,
		languje: languajes[0],
	}

	jesusDS := dataScientist{
		name:    name,
		languje: languajes[2],
	}
	fmt.Println(jesusDev)
	jesusDev.develop()
	speak(jesusDev)
	jesusDS.develop()
	speak(jesusDS)

	//Anonymous function
	// func(parameters){body}(arguments)
	func(name string) {
		println("Anomymous function was called by", name)
	}(name)

	//func expression
	//functios are first class citizens
	//A funtion is also a type
	funcExpression := func(name string) {
		println("funcExpression was called by", name)
	}
	funcExpression(name)

	// returning a function
	y := x()
	fmt.Printf("Type: %T \t Value: %v\n", y, y())
}
