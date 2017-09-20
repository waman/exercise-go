// 【練習問題 4.9】
// 入力テキストファイル内のそれぞれの単語の出現頻度を報告するプログラム
// wordfreq を書きなさい。入力を行ではなく単語へ分割するために、
// 最初の Scan 呼び出しの前に input.Split(bufio.ScanWords) を
// 呼び出しなさい。
package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"log"
)

// 実行例：
//
//   > go run ./ch4/ex4_9/main.go ./ch4/TheGoBlog-strings.txt
//
// go get でコードを取得した場合は、上記のコマンドではうまく動かないかもしれません。
func main(){
	var r io.Reader

	if len(os.Args) <= 1 {
		r = os.Stdin
	}else {
		// 引数があればファイルから読み取る（os パッケージのドキュメント参照）
		file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0400)
		if err != nil { log.Fatal(err) }

		defer func(){
			if cErr := file.Close(); err == nil && cErr != nil {
				log.Fatal(err)
			}}()

		r = file
	}

	counts := wordfreq(r)

	fmt.Printf("\nword\tcount\n")
	for w, n := range counts {
		fmt.Printf("%s\t%d\n", w, n)
	}
}

func wordfreq(r io.Reader) map[string]int {
	counts := make(map[string]int)

	input := bufio.NewScanner(r)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}

	return counts
}

