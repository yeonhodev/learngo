package main

import "fmt"

func main() {
	var a []int = []int{5, 6, 7, 8, 9}
	fmt.Println(cap(a))
	fmt.Println(cap(a[:3]))
	a = append(a, 10)
	fmt.Println(a)

	b := make([]int, 5)
	fmt.Println(b)
	fmt.Printf("%T", b)
}
