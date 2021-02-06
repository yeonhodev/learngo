package main

import (
	"fmt"
)

func main() {
	ans := 10

	switch ans {
	case 1:
		fmt.Println("1. one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not a case")
	}
}
