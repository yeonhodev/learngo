package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	retur false
}

func main() {
	fmt.Println(canIDrink(16))
}
