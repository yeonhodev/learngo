package main

import "fmt"

func main() {
	var a []int = []int{1, 3, 4, 56, 7, 12, 4, 6}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for i, element := range a {
		fmt.Printf("%d: %d\n", i, element)
	}
}
