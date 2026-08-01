// 구조체 심화(5)
package main

import "fmt"

type Employee2 struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee2) Calculate() float64 {
	return e.salary + e.bonus
}

// 이름이 같은 함수 사용 : 오버라이딩
func (e Executives2) Calculate() float64 {
	return e.Employee2.salary + e.Employee2.bonus + e.specialBonus
}

type Executives2 struct {
	Employee2
	specialBonus float64
}

func main() {
	//구조체 임베디드 메소드 오버라이딩 패턴
	//
	//예제1
	//직원
	ep1 := Employee2{"kim", 2000000, 300000}
	ep2 := Employee2{"park", 1500000, 200000}
	//임원
	ex := Executives2{
		Employee2{"lee", 4000000, 1000000},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	fmt.Println("ex1 : ", int(ex.Calculate()+ex.specialBonus)) //오버라이딩 : 잘못 된 값 반환
	fmt.Println("ex1 : ", int(ex.Calculate()))                 //오버라이딩 : 정확한 값
	fmt.Println("ex1 : ", int(ex.Employee2.Calculate()+ex.specialBonus))

}
