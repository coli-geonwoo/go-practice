// 데이터 타입 : Numeric 연산(1)
package main

import "fmt"
import "math"

func main() {
	//숫자 연산(산술, 비교)
	//타입이 같아야 가능
	//다른 타입끼리는 반드시 형 변환 후 연산
	//형 변환 없을 경우 예외(에러) 발생
	// +, -, *, /, %, >>, <<, &, ^..

	//예제1 : 최대값 확인
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("ex1 : ", n1)
	fmt.Println("ex2 : ", n2)
	fmt.Println("ex3 : ", n3)
	fmt.Println("ex4 : ", n4)
	fmt.Println("ex5 : ", math.MaxInt8)
	fmt.Println("ex6 : ", math.MaxInt16)
	fmt.Println("ex7 : ", math.MaxInt32)
	fmt.Println("ex8 : ", math.MaxInt64)
	fmt.Println("ex9 : ", math.MinInt8)
	fmt.Println("ex10 : ", math.MinInt16)
	fmt.Println("ex11 : ", math.MinInt32)
	fmt.Println("ex12 : ", math.MinInt64)
	fmt.Println("ex13 : ", math.MaxFloat32)
	fmt.Println("ex14 : ", math.MaxFloat64)

	//예제2 (형 변환)
	n5 := 100000 //int- 기본값
	n6 := int64(10000)
	n7 := uint8(100)

	//fmt.Println("ex2 : ", n5+n6) //예외 발생(컴파일 에러)
	//fmt.Println("ex2 : ", n6+n7) //예외 발생(컴파일 에러)
	fmt.Println("ex2 : ", n5+int(n6))
	fmt.Println("ex2 : ", n6+int64(n7))
	fmt.Println("ex2 : ", n6 > int64(n7))
	fmt.Println("ex2 : ", n6-int64(n7) > 5000)
}
