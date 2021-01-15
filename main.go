package main

import (
	"fmt"
)

func main() {
	a := 2
	// &a is memory address of a
	b := &a
	// change value of a by using pointer
	*b = 20
	// *b is value of the memory address
	fmt.Println(a)
}
