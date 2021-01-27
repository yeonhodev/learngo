package main

import "fmt"

/* 함수2 - 일급 함수, 익명 함수, 클로저(Closure)
Go 언어의 함수는 일련의 코드를 실행해 주는 함수의 역할 뿐만 아니라,
변수에 할당할 수 있고, 함수의 인자나 리턴 값이 될 수도 있다.
*/

func main() {
	hamsu := func(x int) int {
		return x * -1
	}(7)
	fmt.Println(hamsu)
}
