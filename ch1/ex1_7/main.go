// 【練習問題 1.7】
// 関数呼び出し io.Copy(dst, src) は、src から読み込み dst へ書き込みます。
// ストリーム全体を保持するのに十分な大きさのバッファを要求することなくレスポンスの
// 内容を os.Stdout へコピーするために、ioutil.ReadAll の代わりにその関数を
// 使いなさい。なお、io.Copy のエラー結果は必ず検査するようにしなさい。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
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
		// 読み書きした文字数を表示
		fmt.Printf("%d characters are read", written)
	}
}
