package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")

	if err != nil {
		log.Fatalln("error:", err)
	}

	fmt.Println("Errors")
}
