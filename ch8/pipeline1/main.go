// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import "fmt"

func main(){
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func(){
		for x := 0;; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func(){
		for {
			x := <- naturals
			squares <- x * x
		}
	}()

	// Printer
	for {
		fmt.Println(<- squares)
	}
}
