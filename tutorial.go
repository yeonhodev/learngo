package main

import "fmt"

func main() {
	test := func(x int) int {
		return x * -1
	}(8)
	fmt.Println(test)
}
