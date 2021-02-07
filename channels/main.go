package main

import "fmt"

func main() {
	//Channels blocks
	//biderectional channel
	//unbuffer channel
	c1 := make(chan string)
	//buffered channel
	c2 := make(chan int, 2)
	//receiver
	cr := make(<-chan int)
	//sender
	cs := make(chan<- int)
	go func(str string) {
		c1 <- str
	}("Hola")

	c2 <- 2
	c2 <- 3

	fmt.Println(<-c1)
	fmt.Println(<-c2)
	fmt.Println(<-c2)

	fmt.Printf("%T\t\n", c2)
	fmt.Printf("%T\t\n", cr)
	fmt.Printf("%T\t\n", cs)

	c3 := make(chan int)
	go foo(c3)
	bar(c3)
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
