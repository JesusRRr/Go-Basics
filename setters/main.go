package main

import (
	"fmt"
	"math"
)

// Method set determine what methos attach to a Type
type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

//Non pointer receiver
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}
func main() {
	c := circle{5}
	info(c)
}
