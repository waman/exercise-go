// server1 は、最小限の「echo」サーバです。
package main

import (
	"net/http"
	"log"
	"fmt"
)

func main(){
	http.HandleFunc("/", handler)
	log.Println("Server1 starts...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
