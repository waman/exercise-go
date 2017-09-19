// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// downserver は、レスポンスを返さないサーバです。
// wait がサーバからのレスポンスを受け取れずに失敗することを
// エミュレートするためのサーバです。
package main

import (
  "net/http"
  "log"
	"time"
)

func main(){
	http.HandleFunc("/", handler)
	log.Println("downserver starts...")  // 起動メッセージ
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	time.Sleep(2 * time.Minute)
}

