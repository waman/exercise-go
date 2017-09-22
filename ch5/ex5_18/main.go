//【練習問題 5.8】
// 振る舞いを変えることなく、書き込み可能なファイルを閉じるために defer を
// 使うように fetch を書き直しなさい。
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	file, n, err := fetch(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s の内容を %s へ書き出しました： %d バイト", os.Args[1], file, n)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	// テスト用：上記3行の代わりにエラーを返す Writer/Closer を作成
	//var f = writeCloser{}

	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

// Write と Close でエラーを返す Writer/Closer
type writeCloser struct{}

func (wc writeCloser) Write(bs []byte) (int, error) {
	//fmt.Println(string(bs))
	//return len(bs), nil
	return 0, fmt.Errorf("書き出しに失敗")
}

func (wc writeCloser) Close() error {
	//fmt.Println("[Close]")
	//return nil
	return fmt.Errorf("閉じるのに失敗")
}
