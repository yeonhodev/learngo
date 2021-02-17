package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 1000
}

//mutable data type
func main() {
	var x []int = []int{3, 4, 5}
	y := x
	y[0] = 100
	fmt.Println(x, y)

	var a map[string]int = map[string]int{"hello": 3}
	b := a
	b["b"] = 100
	fmt.Println(a, b)

	var c [2]int = [2]int{3, 4}
	d := c
	d[0] = 100
	fmt.Println(c, d)

	changeFirst(x)
	fmt.Println(x)
}
