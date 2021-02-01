package main

import "fmt"

func main() {
	var out string = fmt.Sprintf("Number: \t %07d is cool", 45)
	fmt.Println(out)
}
