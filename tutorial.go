package main

import "fmt"

func main() {
	// mp := make(map[string]int)
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	fmt.Println(len(mp))

	// Check if a value exists in a map
	val, ok := mp["apple"]
	fmt.Println(val, ok)

	fmt.Println(mp["apple"])
	mp["apple"] = 900
	mp["tim"] = 10
	delete(mp, "pear")
	fmt.Println(mp)
}
