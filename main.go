package main

type person struct {
	name string
	age int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{"nico", 18, favFood}
	fmt.Println(nico)
}
