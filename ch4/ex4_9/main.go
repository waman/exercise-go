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
)

// 実行例
//
//   > go run ./ch4/ex4_9/main.go < ./ch4/TheGoBlog-strings.txt
//
// go get でコードを取得した場合は、上記のコマンドではうまく動かないかもしれません。
func main(){
	counts := wordfreq(os.Stdin)

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

