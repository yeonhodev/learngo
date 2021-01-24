package main

import "fmt"

/*
변수 앞에 '&'를 붙이면 해당 변수의 메모리 주소를 뜻한다.
따라서 포인터형 변수에 대입할 수 있다.
*는 포인터 변수를 역참조(dereference)해서 포인터형 변수에 값을 대입하거나, 가져올 수 있다.
*/

func 손가락봐라(str *string) {
	//손가락봐라 function에 메모리 주소가 전달되면 그것을 역 참조해 결과적으로 손가락 변수에 저장 된 값을 '별이다!'로 변경한다.
	*str = "별이다!"
}

func 발가락봐라(str string) {
	str = "UFO다!"
}

func main() {
	손가락 := "달이다!"
	fmt.Println(손가락)
	//손가락이라는 function에 손가락의 메모리 주소를 전달
	손가락봐라(&손가락)
	fmt.Println(손가락)
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
