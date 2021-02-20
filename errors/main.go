package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	defer foo()

	if err != nil {
		log.Print("error:", err) // this is the same that using os.Exit()
	}

	fmt.Println("Errors")
}

func foo() {
	fmt.Println("When os.Exit() is called, Deferred functions don't run")
}
