// 데이터 타입 : Numeric(3)
package main

import "fmt"

func main() {
	//데이터 타입 : 숫자형
	//실수형(부동소수점)
	//float32(7자리), float64(15자리)

	//예제1
	// 소수점 사용
	var num1 float32 = 0.14
	var num2 float32 = .75647
	var num3 float32 = 442.0378373
	var num4 float32 = 10.0

	// 지수 표기법
	var num5 float32 = 14e6
	var num6 float64 = .156875e+3
	var num7 float64 = 5.32521e-10

	fmt.Println("ex1 : ", num1)
	fmt.Println("ex2 : ", num2)
	fmt.Println("ex3 : ", num3)
	fmt.Println("ex4 : ", num4-0.1)          //주의1
	fmt.Println("ex5 : ", float32(num4-0.1)) //주의2
	fmt.Println("ex6 : ", float64(num4-0.1)) //주의3
	fmt.Println("ex7 : ", num5)
	fmt.Println("ex8 : ", num6)
	fmt.Println("ex9 : ", num7)

}
