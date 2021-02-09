package main

import "fmt"

func main() {
	arr := [3]int{4, 5, 6}
	sum := 0

	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2D)
	fmt.Println(arr2D[0][2])

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
