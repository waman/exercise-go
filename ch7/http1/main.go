package main

import (
	"log"
	"net/http"
	"fmt"
)

// 実行例：
//
//   > go run ./ch7/http1/main.go
//
//   > go build ./ch1/fetch
//   > fetch http://localhost:8000
//
func main(){
	db := database{"shoes": 50, "socks": 5}

	log.Println("http1 starts...")
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request){
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
