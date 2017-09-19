package main

import (
	"log"
	"net/http"
	"fmt"
)

// 実行例：
//
//   > go run ./ch7/http4/main.go
//
//   > go build ./ch1/fetch
//   > fetch http://localhost:8000/list
//   > fetch http://localhost:8000/price?item=socks
//
func main(){
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	log.Println("http4 starts...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request){
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	}
}


func (db database) price(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
