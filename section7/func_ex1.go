// 함수 심화(1) - 가변인자
package main

import "fmt"

func multiply_1(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}

	return tot
}

func sum_1(n ...int) (tot int) {
	tot = 0
	for _, value := range n {
		tot += value
	}
	return
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2 :", value)
	}
}

func main() {
	//함수 고급
	//가변 인자 실습(매개변수 개수가 동적으로 변할 때 - 정해져있지 않음)

	//예제1
	x := multiply_1(5, 6, 7, 8, 9, 10)
	y := sum_1(5, 6, 7, 8, 9, 10)
	fmt.Println("ex1 :", x)
	fmt.Println("ex1 :", y)
	fmt.Println()

	//예제2
	prtWord("a", "apple", "test", "seoul", "golang", "hi")
	fmt.Println()

	//예제3
	a := []int{5, 6, 7, 8, 9, 10}

	m := multiply_1(a...) //배열... 으로 가변인자 사용 가능
	n := sum_1(a...)

	fmt.Println("ex3 :", m)
	fmt.Println("ex3 :", n)
}
