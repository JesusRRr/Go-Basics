package main

import "fmt"

func main() {
	// Arrays
	// An array is a building block for Slices in go
	// Array declaration
	// the lenght is part of the array's type
	var myArray [6]int
	fmt.Printf("%T Type: %v \n", myArray, myArray)

	//Composite literal
	// Construct values for structs, arrays, slices and maps
	// Is an expression that creates a new instance each time it is evaluated
	// Slice
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T Type: %v \n", mySlice, mySlice)
	//for range through mySlice
	for index, value := range mySlice {
		fmt.Println("index: ", index, "\tvalue:", value)
	}
	//Slicing a slice

	fmt.Println(mySlice[:2])
	fmt.Println(mySlice[2:])
	fmt.Println(mySlice[:3][:2][:1])

	// built-in funcion append
	// Append values to a slice
	mySlice2 := []int{6, 7, 8, 9, 10}
	mySlice2 = append(mySlice, mySlice2...)
	fmt.Println(mySlice2)

	// Deleting from a Slice: deleting the
	mySlice2 = append(mySlice2[:4], mySlice2[5:]...)
	fmt.Println(mySlice2)

	// built-in funcion make
	// make([]T, len, cap)
	mySlice3 := make([]int, 5, 5)
	fmt.Println(mySlice3)
	fmt.Println("len: ", len(mySlice3))
	fmt.Println("cap: ", cap(mySlice3))

	mySlice3 = append(mySlice3, 1)
	fmt.Println("len: ", len(mySlice3))
	fmt.Println("cap: ", cap(mySlice3))

	// multi-dimensional arrays
	muliSlice := [][]int{mySlice3, mySlice3, mySlice3}
	fmt.Printf("%T%v\n", muliSlice, muliSlice)

	// Maps
	//Composite literal
	myMap := map[string]int{
		"jesus": 2,
	}
	fmt.Println(myMap["jesus"])

	// Comma OK idiom
	v, ok := myMap["go"]
	fmt.Println(v)
	fmt.Println(ok)

}
