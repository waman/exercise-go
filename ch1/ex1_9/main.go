// 【練習問題 1.9】
// fetch を修正して、resp.Status に設定されている HTTP ステータスコードも
// 表示するようにしなさい。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
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
		fmt.Printf("[Status Code] %s\n", resp.Status)
		fmt.Printf("[Read Characters] %d characters are read", written)
	}
}
