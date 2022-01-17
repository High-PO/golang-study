package main // 패키지는 메인

import "fmt" // 출력을 위한 모듈

func main() {
	var v1 bool = true  // bool변수를 선언 즉 참이냐 거짓이냐를 불 변수에 넣음
	var v2 bool = false // bool변수를 입력 위에서는 참 여기는 거짓을 넣음

	fmt.Println(v1 == v2) // 두개의 bool 변수를 비교했을때 다르기 때문에 false를 출력
}
