package main

func main() {
	// 잘못된 타입 변수 사용
	// Go는 강타입 언어 중에서도 타입 검사를 가장 엄격하게 합니다.
	a := 3
	var b float64 = 3.5

	var c int = b // 타입이 다른 변수 대입 불가
	d := a * b    // 타입이 다른 변수들 끼리 연산 불가

	var e int64 = 7
	f := a * e // 정수 값이더라도 엄연히 다른 타입이기 때문에 연산 불가
}
