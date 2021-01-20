package main

import "fmt"

func main() {
	a := 3
	// assign memory address of a to b.
	b := &a
	// below is same as fmt.Println(a)
	fmt.Println(*&a)
	// update a's value by using *b.
	*b = 20
	// a is eqal to *b
	fmt.Println(a)
}
