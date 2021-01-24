package main

import "fmt"

/*
변수 앞에 '&'를 붙이면 해당 변수의 메모리 주소를 뜻한다.
따라서 포인터형 변수에 대입할 수 있다.
*는 포인터 변수를 역참조(dereference)해서 포인터형 변수에 값을 대입하거나, 가져올 수 있다.
*/

func main() {
	이것은진리 := "달이다!"
	var 손가락 *string = &이것은진리
	fmt.Println(*손가락, 손가락, &손가락)
}

/*
func main() {
	a := 3
	// assign memory address of a to b.
	b := &a
	// below is same as fmt.Println(a)
	fmt.Println(*&a)
	// update a's value by using *b.
	*b = 20
	// a is eqal to *b
	fmt.Println(a)
}
*/
