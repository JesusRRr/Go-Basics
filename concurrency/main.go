package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("CPU's:", runtime.NumCPU())
	fmt.Println("Go rutines:", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	fmt.Println("CPU's:", runtime.NumCPU())
	fmt.Println("Go rutines:", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}
