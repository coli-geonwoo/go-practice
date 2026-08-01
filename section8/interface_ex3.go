// 인터페이스 고급(3)
package main

import "fmt"
import "reflect"

func main() {
	//타입 변환(Type Assertion)
	//실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	//인터페이스.(타입) 형식 -> 형 변환(인터페이스를 해당 타입으로 변환)
	//interfaceVal.(type)

	//예제1
	var a interface{} = 15.0

	b := a
	//c := a.(int)
	c := a.(float64) //예외 발생 -원래 값이 int였기 때문

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", reflect.TypeOf(a))

	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", reflect.TypeOf(b))

	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", reflect.TypeOf(c))

	fmt.Println()

	//예제2(저장된 타입 실제 타입 검사)
	//v : 원본 값 - 15
	//ok : 형변환 가능한지 출력
	if v, ok := a.(int); ok { //해당 타입 값, 타입 체크 결과
		fmt.Println("ex2 : ", v, ok)
	} else {
		fmt.Println("Can not Type Assertion")
	}

}
