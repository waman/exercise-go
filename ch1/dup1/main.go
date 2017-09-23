// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// dup1 は標準入力から2回以上現れる行を出現回数と一緒に表示します。
package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"log"
)

// 実行例：
//
//   > go run ./ch1/dup1/main.go ./resources/commitors1.txt
//
// commitors1.txt ファイルは、golang の GitHub から、適当な日のコミット実行者を拝借してリストアップしたものです。
// また、go get でコードを取得した場合は、上記のコマンドではうまくいかないかもしれません。
func main() {
	var r io.Reader
	if len(os.Args) == 1 {
		r = os.Stdin
	}else{
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		r = f
	}

	counts := make(map[string]int)
	input := bufio.NewScanner(r)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
