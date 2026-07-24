package main

import "fmt"

func main() {

	// 짧은 선언
	// 함수 안에서만 사용(전역 X), 선언 후 할당하면 예외 발생
	// 재할당하면 선언이 불가함
	// 주로 제한된 범위 내의 함수 내에서 사용할 경우 코드 가독성을 높일 수 있음

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	//shortVar1 := 10 <- 예외 발생

	fmt.Println("shortVar1 : ", shortVar1, "shortVar2 : ", shortVar2, "shortVar3 : ", shortVar3)

	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Succes!")
	}
}
