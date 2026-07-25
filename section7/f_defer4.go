// 함수 Defer(4)
package main

import "fmt"

func start(t string) string {
	fmt.Println("start:", t)
	return t
}
func end(t string) {
	fmt.Println("end:", t)
}

func a() {
	//defer는 가장 부모 함수(파라미터의 함수는 호출됨) start -> in -> defer(end)
	defer end(start("b"))
	fmt.Println("in a")
}

func main() {

	//예제1
	a()
}
