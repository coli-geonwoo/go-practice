// 함수 Defer(2)
package main

import "fmt"

func sayHello(msg string) {
	//가장 마지막에 실행
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi ")
	}()
}

func main() {

	//예제1
	sayHello("Golang!")
}
