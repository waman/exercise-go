// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import "fmt"

func main(){
	f(3)
}

func f(x int){
	fmt.Printf("f(%d)\n", x+0/x)  // x == 0 ならパニック
	defer fmt.Printf("defer %d\n", x)
	f(x-1)
}
