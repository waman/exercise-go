package main

import (
	"time"
	"log"
)

func main(){
	bigSlowOperation()
}

func bigSlowOperation(){
	defer trace("bigSlowOperation")()  // 丸括弧を忘れないように！
	time.Sleep(10*time.Second)  // 遅い操作
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func(){
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
