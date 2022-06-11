package main

import "fmt"

func main() {
	var a int = 3
	var b int
	var c = 4 // type 생략, 우변의 타입으로 지정
	d := 5    // 선언 대입문 : var 와 type 생략
	fmt.Println(a, b, c, d)

	var float32DefaultVal float32
	fmt.Println(float32DefaultVal) // 0.0

	var booleanDefaultVal bool
	fmt.Println(booleanDefaultVal) // false

	var stringDefaultVal string
	fmt.Println(stringDefaultVal) // ""
}
