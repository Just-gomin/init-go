package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 동시 작업이 완료될 때까지 대기

func SumAtoB(a, b int) {
	sum := 0

	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Printf("%d부터 %d까지의 합계는 %d입니다.\n", a, b, sum)
	wg.Done()
}

func main() {
	num := 10
	wg.Add(num) // 대기할 작업의 수 특정

	for i := 0; i < num+2; i++ {
		go SumAtoB(i, 1000000000)
	}

	wg.Wait() // 정해진 수의 작업이 완료될 때까지 대기
}
