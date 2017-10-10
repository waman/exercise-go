// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// パッケージ bank は一つの講座を持つ並行的に安全な銀行を提供します。
package bank1

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int){ deposits <- amount }
func Balance() int      { return <-balances }

func teller(){
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init(){
	go teller()
}