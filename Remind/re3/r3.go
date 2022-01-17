package main // 패키지는 메인

import "fmt" // 출력을 위한 모듈

func main() {
	s1 := "Hello" // s1에다가 Hello라는 문자열 입력
	s2 := "Hello" // s2에다가 Hello라는 같은 문자열 입력
	s3 := "World" // s3에다가 World라는 문자열 입력

	fmt.Println(s1 == s2) // s1과 s2를 비교한다. 비교해서 같으면 true 다르면 false 출력
	fmt.Println(s1 + s3)  // s1과 s3를 더한다. 즉 Hello와 World를 더해서 HelloWorld 출력
}
