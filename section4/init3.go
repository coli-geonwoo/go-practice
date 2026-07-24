// GO 초기화 함수(3)
package main

import (
	"fmt"
	"untitled/section4/lib"
)

var num int32

// 변수 초기화
func init() {
	num = 30
}

func main() {
	fmt.Print("10 보다 큰 수? :  ", num, lib.CheckNum1(num))
}
