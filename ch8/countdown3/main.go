// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"fmt"
	"time"
	"os"
)

func main(){
	abort := make(chan struct{})
	go func(){
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// 何もしない
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}
	launch()
}

func launch(){
	fmt.Println("Launch!")
}
