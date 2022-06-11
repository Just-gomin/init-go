package main

import "fmt"

func main(){
	// var keyword를 이용해 변수 선언을 시작
	// var - 변수 명 - 변수 타입 - [= 값] 을 통해 메모리 할당
	var a int = 10
	var msg string = "Hello Variable"

	// 변수의 값 변경
	a = 20
	msg = "Good Morning"

	// 변수를 이용해 값을 전달
	fmt.Println(msg, a)
}