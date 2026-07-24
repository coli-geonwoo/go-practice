// 라이브러리 및 접근제어(1)
package lib

import "fmt"

func CheckNum1(c int32) bool {
	return c > 10
}

func init() {
	fmt.Println("lib 로드 완료!")
}
