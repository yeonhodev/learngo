package main

import (
	"fmt"
)

func main() {
	ans := 10

	switch {
	case ans > 0:
		fmt.Println("greater than 0")
	case ans < 0:
		fmt.Println("less than 0")
	default:
		fmt.Println("zero")
	}
}
