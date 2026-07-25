// 자료형 : 맵(4)
package main

import "fmt"

func main() {
	//맵(Map)
	//맵 조회 할 경우 주의 할점
	//value, isExists = map[key]로 존재하는지에 대한 boolean 값을 같이 받을 수 있음

	//예제1
	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, isExists := map1["kiwi"] //value, 키가 있는지

	fmt.Println("ex1 : ", value1)
	fmt.Println("ex1 : ", value2)
	fmt.Println("ex1 : ", value3, isExists) //두 번째 리턴 값으로 키 존재 유무 확인
	fmt.Println()

	//예제2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist!")
	}

	if value, ok := map1["lemon"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : lemon is not exist!")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2 : kiwi is not exist!")
	}
}
