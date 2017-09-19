// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// server2 は最小限の「echo」とカウンタ機能を持つサーバです。
package main

import (
	"net/http"
	"sync"
	"log"
	"fmt"
)

var mu sync.Mutex
var count int

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Println("Server2 starts...")  // 起動メッセージ
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
