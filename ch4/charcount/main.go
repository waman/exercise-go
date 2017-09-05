// charcount は Unicode 文字の数を計算します。
package main

import (
	"unicode/utf8"
	"bufio"
	"os"
	"io"
	"fmt"
	"unicode"
)

// 実行例
//
//   > go run ./ch4/charcount/main.go < ./ch4/TheGoBlog-strings.txt
//
// ch4 ディレクトリ下に TheGoBlog-strings.txt を置いていますが、
// これは The Go Blog の記事 "https://blog.golang.org/strings"
// から拝借しました（いくつか漢字などが含まれている、ちょうどよさげな内容だったため）。
// go get でコードを取得した場合は、上記のコマンドではうまく動かないかもしれません。
func main(){
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
