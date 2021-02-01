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

type personJS struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
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

	fmt.Print(string(bs))
	fmt.Printf("\t%T\n", bs)
	peopleJS := []personJS{}
	//func UnMarshal(data []byte, v interface{}) error
	//Unmarshal parses the JSON-encoded data and stores the value pointed to v
	err = json.Unmarshal(bs, &peopleJS)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(peopleJS)
	fmt.Printf("\t%T\n", peopleJS)
}
