package main

import "fmt"

func main() {
	//Channels blocks

	//unbuffer channel
	c1 := make(chan string)
	//buffered channel
	c2 := make(chan int, 2)

	go func(str string) {
		c1 <- str
	}("Hola")

	c2 <- 2
	c2 <- 3

	fmt.Println(<-c1)
	fmt.Println(<-c2)
	fmt.Println(<-c2)
}
