package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) setAge(age int) {
	s.age = age
}

func main() {
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	fmt.Println(s1)
	s1.setAge(7)
	fmt.Print(s1)
}
