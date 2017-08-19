// 【練習問題 1.8】
// fetch を修正して、個々の引数の URL に接頭辞 http:// がなければ
// 追加するようにしなさい。strings.HasPrefix を使いたくなるかもしれません。
package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
)

func main(){
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		written, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: Reading %s: %v\n", url, err)
		}
		fmt.Printf("%d characters are read", written)
	}
}
