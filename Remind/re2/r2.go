package main // 패키지는 왠만해선 메인으로 하기

import (
	"fmt"          // fmt는 println을 위한 모듈
	"unicode/utf8" // unicode(한글)을 인식시키기 위한 모듈
)

func main() {
	s1 := "Hello"        // 간단하게 s1에다가 Hello라는 문자열 추가를 위해서 사용
	var s2 string = "한글" // var로 된 변수만 수정이 가능한데 변수 수정을 위해서 var s2 string = "한글"을 사용함

	fmt.Println(len(s1))                    // s1에 대한 길이를 알려줌(영어)
	fmt.Println(utf8.RuneCountInString(s2)) //s2에 대한 길이를 알려줌(한글)
	fmt.Println(s2)                         // s2 내용 출력

	s2 = "하이여 ㅋ"    // s2의 내용 변경
	fmt.Println(s2) // 변경된 s2 내용 출력
}
