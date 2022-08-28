package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("[ch] Square: %d\n", n*n)
			time.Sleep(time.Second)
		case run := <-quit:
			fmt.Printf("[quit] Quit: %t\n", run)
			if run {
				fmt.Println("In Quit Proc")
				wg.Done()
				return
			} else {
				fmt.Println("In Proc")
				time.Sleep(time.Second)
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- false
	quit <- true
	wg.Wait()
}
