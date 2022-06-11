package main

import "fmt"

func main() {
	// 타입 변환
	// 다른 타입으로 변환하고자 하는 경우,
	// 타입명 뒤에 괄호를 쓰고 그안에 변수를 넣음으로써 타입 변환이 가능합니다.
	a := 3
	var b float64 = 3.5

	var c int = int(b)  // float64 3.5 -> int 3 :: 실수형에서 정수형으로의 변화시에는 소수점아래의 값을 "버림" 합니다.
	d := float64(a * c) // int -> float64

	var e int64 = 7
	f := int64(d) * e // float64 -> int64

	var g int = int(b * 3) // float64 -> int로 형 변환 전에 연산을 진행해, 3.5 * 3 = 10.5가 된 후 형 변환을 진행
	var h int = int(b) * 3 // float64 -> int로 형 변환을 선행 후 연산을 진행 해 3 * 3 = 9로 연산
	fmt.Println(g, h, f)
}
