package main

import "fmt"

// 함수
// 함수는 특정 기능을 위해 만든 여러 문자을 묶어서 실행하는 코드 블럭 단위이다.
// 기본적인 형태의 함수 선언은 "func 함수이름 (매개변수이름 매개변수형) 반환형"이다.
// 매개변수형 = 데이터 타입
// 매개변수(parameter), 전달인자(argument) = 호출할 때 함수 안에 넣어주는 숫자 5, 7

func add(x, y int) (w1, w2 int) {
	defer fmt.Println("나는 반환이 끝난 후 실행되요.")
	w1 = x + y
	w2 = x - y
	fmt.Println("전 반환 작업 중입니다.")
	return
}

func main() {
	z1, z2 := add(5, 7)
	fmt.Println(z1, z2)
}
