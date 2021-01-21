package main

import "fmt"

func main() {
	// A missing switch expression is equivalent to true
	switch {
	case false:
		fmt.Println("Should no print this")
	case (5 == 6):
		fmt.Println("Should no print this")
	case true:
		fmt.Println("Should print this")
	}

	// cases can be presented in comma-separated list
	day := "Tuesday"
	switch day {
	case "Monday", "Tuesday", "Wednesday ", "Thursday", "Friday":
		fmt.Println(day, " is a Week day")
	case "Saturday", "Sunday":
		fmt.Println(day, " is a Weekend day")
	}
}
