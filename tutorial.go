package main

import "fmt"

type Point struct {
	x int32
	y int32
}

func main() {
	var p1 Point = Point{1, 2}
	//var p2 Point = Point{-5, 7}
	fmt.Println(p1.x)
}
