package main

import "fmt"

func main() {
	a := 2
	b := 5
	// Put "&" in front of variables to see memory address.
	fmt.Println(&a, &b)
}
