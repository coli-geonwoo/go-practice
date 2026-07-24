package main

import "fmt"

func main() {
	//열거형
	//상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 활동
	const (
		Jan = iota + 1
		Feb
		Mar
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)

}
