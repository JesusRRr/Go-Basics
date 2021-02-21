package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	defer foo()
	defer bar()

	if err != nil {
		//log.Fatal("error:", err) this is the same that using os.Exit()
		//log.Println("error:", err) this is the same that using os.Exit()
		panic(err) // this stop the current go rutine but the deferred functions will run
	}

	fmt.Println("Errors")
}

func foo() {
	fmt.Println("When os.Exit() is called, Deferred functions don't run")
}

func bar() {
	fmt.Println("bar function deffered")
}
