// 함수 Defer(3)
package main

import "fmt"

func stack() {
	//하나의 함수에서 defer는 stack처럼 쌓임 -> LIFO
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex1 : ", i)
	}
}

func main() {

	//예제1
	stack()
}
