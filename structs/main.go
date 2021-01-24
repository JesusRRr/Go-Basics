package main

import "fmt"

// Struct
// A struct is a composite data type
// since compose together values of diferent type
type person struct {
	name     string
	lastName string
}

type languaje struct {
	name  string
	level int
}

type developer struct {
	person
	languajes []languaje
}

func main() {
	// Composite literal
	// value of type person assigned to variable jesus
	jesus := person{
		name:     "Jesus",
		lastName: "R",
	}
	fmt.Printf("Type: %T\nValue: %v\n\n", jesus, jesus)

	languajes := []languaje{
		{
			name:  "go",
			level: 4,
		},
		{
			name:  "java",
			level: 6,
		},
		{
			name:  "python",
			level: 4,
		},
		{
			name:  "javascript",
			level: 3,
		},
	}

	// Embedded struct
	goDeveloper := developer{
		person: jesus,

		languajes: languajes,
	}
	fmt.Printf("Type: %T\nValue: %v\n\n", goDeveloper, goDeveloper)

	fmt.Printf("Type: %T\nValue: %v\n\n", languajes, languajes)

	// Anonymous struct
	file := struct {
		code string
	}{
		code: "main.go",
	}

	fmt.Printf("Type: %T\nValue: %v\n\n", file, file)
}
