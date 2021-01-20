package main

import "fmt"

func main() {
	a := 2
	// Save a's memory addres in b.
	b := &a
	// Put "&" in front of variables to see memory address.
	fmt.Println(&a, b)
}
