package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '비', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}

func main() {
	go PrintHangul()  // 한글 출력 고 루틴 생성
	go PrintNumbers() // 숫자 출력 고 루틴 생성

	// time.Sleep(3 * time.Second)	// 3초 대기
}
