// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package bank3

import "sync"

var(
	mu sync.Mutex
	balance int
)

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func Deposit(amount int){
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int{
	mu.Lock()
	defer mu.Unlock()
	return balance
}

// この関数は、ロックが獲得されていることを前提としている。
func deposit(amount int){ balance += amount }