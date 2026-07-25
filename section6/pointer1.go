// 자료형 : 포인터(1)
package main

import "fmt"

func main() {
	//포인터
	//Go : 포인터 지원(C, C++)
	//포인터지원 X(파이썬, C#, JAVA 등)
	//주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
	//메모리 주소를 출력(값의 메모리 주소)
	//애스터리스크로 사용(*)
	//nil로 초기화(0 아님)

	//예제1
	var a *int            //방법1
	var b *int = new(int) //방법2

	fmt.Println(a) //nil
	fmt.Println(b) //주소값을 가지고 있음

	i := 7
	a = &i //& 주소값 전달
	b = &i

	fmt.Println("*a : ", *a)
	fmt.Println("*b : ", *b)

	var c = &i //방법3
	d := &i    //방법4

	fmt.Println("ex1 : ", i)
	fmt.Println("ex1 : ", &i) //실행 할 때 마다 시스템 별로 변경

	fmt.Println("ex1 : ", *a)
	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", &a) //포인터 변수 a의 주소 값

	fmt.Println("ex1 : ", *b)
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", &b) //포인터 변수 b의 주소 값

	fmt.Println("ex1 : ", *c)
	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", &c) //포인터 변수 c의 주소 값

	fmt.Println("ex1 : ", *d)
	fmt.Println("ex1 : ", d)
	fmt.Println("ex1 : ", &d) //포인터 변수 d의 주소 값
}
