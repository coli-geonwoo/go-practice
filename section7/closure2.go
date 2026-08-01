// 함수 Closure(2)
package main

import "fmt"

func main() {

	//예제1
	cnt := increaseCnt() //n이 캡처된 함수를 cnt가 할당 받게 됨

	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())
	fmt.Println("ex1 : ", cnt())

}

func increaseCnt() func() int {
	n := 0 //지역변수(캡처됨)
	return func() int {
		n += 1
		return n
	}
}
