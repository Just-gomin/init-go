package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex // 자원 접근 통제

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := Account{0}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(&account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()         // 타 작업의 자원 접근을 막음
	defer mutex.Unlock() // 접근 해제

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	fmt.Printf("Balance: %d\n", account.Balance)
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
