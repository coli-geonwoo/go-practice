// 자료형 : 슬라이스(1)
package main

import "fmt"

func main() {
	//슬라이스(길이 & 용량 개념)
	//슬라이스 배열과 같지만 크기가 동적 할당 가능
	//배열 vs 슬라이스 차이점 중요
	//길이 고정 						   vs 길이 가변
	//값 타입 								  vs 참조 타입(레퍼런스)
	//복사 전달 							 vs  참조 값 전달
	//전체 비교연산자 사용 가능 vs 비교 연산자 사용 불가
	//대 부분 슬라이스 사용한다.

	//make(자료형, 길이, 용량(생략시 길이)) : 할당해야 값 삽입 가능

	//예제1
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10, 11}, //콤마 주의
	}
	//slice2[0] = 1 예외 발생(가변형이기 때문에 아직 초기화되지 않은 값에 대해 수정 불가)
	//slice3[4] = 10 //값 수정 가능

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)
	fmt.Println()

	//예제2
	var slice5 []int = make([]int, 5)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 10)
	slice8 := make([]int, 5)
	slice6[2] = 7 //삽입

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8) //길이와 용량 같다.
	fmt.Println()

	//예제3
	var slice9 []int //nil 슬라이스 (길이와 용량 0)
	if slice9 == nil {
		fmt.Println("This is Nil Slice.")
	}
}
