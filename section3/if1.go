package main

import "fmt"

func main() {

	//조건문
	//if문은 반드시 Boolean으로 검사 -> 1,0으로 판단 안됨
	//소괄호를 사용하지 않음

	var a int = 20
	b := 20

	//예제1
	if a >= 15 {
		fmt.Println("15이상")
	}

	//예제2
	if b >= 25 {
		fmt.Println("25이상")
	}

	//에러 발생1 -> 컴파일하면서 뒤에 ;를 붙여서 에러남
	//if b>=25
	//{
	//
	//}

	//에러 발생2 - 중괄호가 없는 경우
	// if b>=25
	// fmt.Println

	//에러 발생3 - Booelan이 아닌 값을 조건문에 넣으려 할때
	if c := true; c {
		fmt.Println("True")
	}

	//예제
	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}

	// c+=20 짧은 선언이라 이미 소멸되었으므로 에러 남
}
