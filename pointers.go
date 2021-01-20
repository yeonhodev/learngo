package main

import "fmt"

func main() {
	a := 3
	// Save a's memory addres in b.
	// Put "&" in front of variables to see memory address.
	b := &a
	// Put "*" in front of a variable to get a value of a memory address.
	fmt.Println(b, *b)
}
