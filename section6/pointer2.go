// 자료형 : 포인터(2)
package main

import "fmt"

func main() {

	//예제1
	type bag struct{ witdh, height, weight float32 } //구조체
	var p *int = new(int)
	var p_bag *bag = &bag{20, 50, 30}

	fmt.Println("ex1 : ", p)
	fmt.Println("ex1 : ", p_bag)
	fmt.Println()

	//p++              //컴파일 에러, 포인터 연산 허용 X
	//p = 0xc071405232 ////컴파일 에러, 주소값 대입 허용 X

	//예제2
	i := 7
	p = &i

	fmt.Println("ex2-1 : ", i, *p, &i, p)

	*p++ //1 증가
	fmt.Println("ex2-2: ", i, *p, &i, p)

	*p = 10 //포인터 변수 역 참조 값 변경
	fmt.Println("ex2-3 : ", i, *p, &i, p)

	i = 77 //원본 변수 값 변경
	fmt.Println("ex2-4 : ", i, *p, &i, p)

}
