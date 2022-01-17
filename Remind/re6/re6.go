package main

import "fmt"

func main() {
	/*
		const a = 0
		const b = 1
		const c = 2
		const d = 3
		const e = 4
		const f = 5
		위와 같이 작성하면 너무 번거로움
	*/

	// ---------------------------------------------

	/*
		const (
			a = 0
			b = 1
			c = 2
			d = 3
			e = 4
			f = 5
		)
		fmt.Println(a)
		괄호를 사용해서 const를 간단하게 선언
	*/

	// -----------------------------------------------

	const (
		a = iota // 0
		b        // 1
		c        // 2
		d        // 3
		e        // 4
		f        // 5
	)
	fmt.Println(d) //3을 출력
	// iota는 자동으로 0부터 순서대로 값을 저장해줌
}
