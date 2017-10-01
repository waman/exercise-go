// 【練習問題 1.10】
// 大量のデータを生成するウェブサイトを見つけなさい。
// 報告される時間が大きく変化するかを調べるために fetchall を
// 2回続けて実行して、キャッシュされているかどうかを調査しなさい。
// 毎回同じ内容を得ているでしょうか。fetchall を修正して、
// その出力をファイルへ保存するようにして調べられるようにしなさい。
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// fetchall と異なり、一度に1つのサイトしか取得できません。
// 実行例：
//
//   > go build ./ch1/ex1_10/main.go
//   > ex1_10 http://amazon.co.jp amazon1.html
//   > ex1_10 http://amazon.co.jp amazon2.html
//
func main() {
	var w io.Writer

	switch len(os.Args) {
	case 1:
		fmt.Println("取得するサイトと、必要なら出力するファイル名を指定してください。")
	case 2:
		w = os.Stdout
	default:
		// 第2引数があればファイルに出力（os パッケージのドキュメント参照）
		file, err := os.Create(os.Args[2])
		if err != nil { log.Fatal(err) }

		defer func() {
			if err := file.Close(); err != nil { log.Fatal(err) }
		}()

		w = file
	}

	start := time.Now()
	ch := make(chan string)
	go fetch(os.Args[1], w, ch)
	fmt.Println(<-ch)

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, w io.Writer, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(w, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
