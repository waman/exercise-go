package bank3

import (
	"fmt"
)

func ExampleBank1(){
	Deposit(100)
	Deposit(200)
	fmt.Println(Balance())
	// Output:
	// 300
}
