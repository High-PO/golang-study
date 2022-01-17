package main // 패키지는 메인

import "fmt" // 출력을 위한 패키지

func main() {
	var s1 string = "Edit"       // 변수 s1을 선언(string)
	const s2 string = "Non Edit" // 상수 s2를 선언 string, int같은 것도 안됨

	fmt.Println(s1) // s1을 출력
	fmt.Println(s2) // s2를 출력

	s1 = "Edit OK" // s1내용 변경
	//s2 = "Error"   // s2 내용 변경(오류가 뜨는게 맞음 상수는 변경이 안됨)

	fmt.Println(s1) // s1 변경된 내용 출력
	fmt.Println(s2) // s2는 변경 할 수 없기에 그대로 출력

}
