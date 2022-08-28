package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// contextWithCancelProc()
	// contextWithTimeOutProc()
	// contextWithValueProc()
	contextWithCancelAndValueProc()
}

// 1. 취소 가능한 Context
func contextWithCancelProc() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) // 취소 가능한 컨텍스트
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()
}

// 2. 작업 시간 제한을 둔 Context
func contextWithTimeOutProc() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go PrintEverySecond(ctx)
	time.Sleep(8 * time.Second)
	cancel()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context Done")
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}

// 3. 특정 값을 설정한 컨텍스트
func contextWithValueProc() {
	wg.Add(1)
	ctx := context.WithValue(context.Background(), "number", 9)
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square: %d\n", n*n)
	}
	wg.Done()
}

// 4. 취소 가능하며 값도 설정하는 컨텍스트
func contextWithCancelAndValueProc() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "number", 9)
	ctx = context.WithValue(ctx, "name", "Brown")

	printContextValue(ctx)

	time.Sleep(4 * time.Second)
	cancel()
}

func printContextValue(ctx context.Context) {
	if number := ctx.Value("number"); number != nil {
		number = number.(int)
		fmt.Printf("number: %d\n", number)
	}

	if name := ctx.Value("name"); name != nil {
		fmt.Printf("name: %s\n", name)
	}

	wg.Done()
}
