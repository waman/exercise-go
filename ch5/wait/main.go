// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"time"
	"net/http"
	"log"
	"fmt"
	"os"
)

// 実行例（下記2つを別のコマンドラインから実行）：
//
//   > go run ./ch5/downserver/main.go
//   > go run ./ch5/wait/main.go http://localhost:8080
//
func main(){
	//log.SetPrefix("wait: ")
	//log.SetFlags(0)

	if err := WaitForServer(os.Args[1]); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}

// WaitForServer は URL のサーバーへ接続を試みます。
// 指数バックオフを使って一分間試みます。
// すべての試みが失敗したらエラーを報告します。
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil  // 成功
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))  // 指数バックオフ
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}