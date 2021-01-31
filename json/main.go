package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		"Jesus",
		"R",
		24,
	}
	p2 := person{
		"Koy",
		"N",
		25,
	}

	people := []person{p1, p2}
	//func Marshal(v interface{}) ([]byte, error)
	// Marshal returns the JSON encoding of v
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
