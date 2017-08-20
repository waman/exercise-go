// 【練習問題 1.10】
// 大量のデータを生成するウェブサイトを見つけなさい。
// 報告される時間が大きく変化するかを調べるために fetchall を
// 2回続けて実行して、キャッシュされているかどうかを調査しなさい。
// 毎回同じ内容を得ているでしょうか。fetchall を修正して、
// その出力をファイルへ保存するようにして調べられるようにしなさい。
package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
)

func main(){
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch10(url, ch)
	}
	for range os.Args[1:]{
		fmt.Println(<- ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch10(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// ファイルへ直接書き出す方法はまだ出てきていないので
	// 標準出力へ書き出して、実行時にファイルを指定して送る。
	//   $ go build exercise10.go
	//   $ exercise10 http://amazon.co.jp > amazon1.html
	//   $ exercise10 http://amazon.co.jp > amazon2.html
	nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
