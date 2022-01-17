package main // 패키지 이름은 main이고 단 하나밖에 존재 할 수 없다.

import "fmt" // 모듈 fmt 선언

func re1() {
	var s1 string = "HelloWorld" // var는 변수 선언 이다. string은 문자열이라는 뜻이다.
	fmt.Println(s1)              // fmt라는 모듈을 사용해서 출력한다.
	// `` <- 해당 백쿼트를 사용할 경우 문자열의 여러줄을 입력 할 수 있다.
	s2 := `안녕하세요.  
	Hello World` // :=의 뜻은 위에서 사용한 s1 string = 과 비슷한 느낌이다. 그냥 s2 안에다가 := 뒤 내용을 넣는다는 뜻이다.
	// 단, func 안에서만 사용이 가능하다. 밖에서는 var를 사용해야한다.
	fmt.Println(s2) // fmt라는 모듈을 통해서 s2를 출력했다.
}
