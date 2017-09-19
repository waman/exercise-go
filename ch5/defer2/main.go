// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"runtime"
	"fmt"
	"os"
)

func main(){
	defer printStack()
	f(3)
}

func printStack(){
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

// defer1 の f と同じ。
func f(x int){
	fmt.Printf("f(%d)\n", x+0/x)  // x == 0 ならパニック
	defer fmt.Printf("defer %d\n", x)
	f(x-1)
}


