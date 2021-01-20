package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10) // iota value=1 10
	mb = 1 << (iota * 10) // iota value=2 100
	gb = 1 << (iota * 10) // iota value=3 1000
)

func main() {
	fmt.Printf("KB -> %b bits \t\t\t\t%d bytes\n", kb, kb)
	fmt.Printf("MB -> %b bits \t\t%d bytes\n", mb, mb)
	fmt.Printf("GB -> %b bits \t%d bytes\n", gb, gb)
}
