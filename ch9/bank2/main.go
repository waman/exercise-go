// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package bank2

var(
	sema = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int){
	sema <- struct{}{}
	balance = balance + amount
	<- sema
}

func Balance() int{
	sema <- struct{}{}
	b := balance
	<- sema
	return b
}