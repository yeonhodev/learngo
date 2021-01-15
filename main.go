package main

import (
	"fmt"
)

func main() {
	a := 2
	// &a is memory address of a
	b := &a
	// *b is value of the memory address
	fmt.Println(*b)
}
