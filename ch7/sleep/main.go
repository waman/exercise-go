package main

import (
	"flag"
	"time"
	"fmt"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

// 実行例：
//
//   > go build ./ch7/sleep
//   > sleep -period 50ms
//   > sleep -period 1m30s
//
func main(){
	flag.Parse()
	fmt.Printf("Sleeping fo %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
