// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// dup2 は入力に2回以上現れた行の数とその行のテキストを表示します。
// 標準入力から読み込むか、名前が指定されたファイルの一覧から読み込みます。
package main

import (
	"os"
	"fmt"
	"bufio"
)

// 実行例：
//
//   > cd ch1
//   > go run ./dup2/main.go commitors1.txt commitors2.txt
//
// commitors#.txt ファイルは、golang の GitHub から、適当な日のコミット実行者を拝借してリストアップしたものです。
// また、go get でコードを取得した場合は、上記のコマンドではうまくいかないかもしれません。
func main(){
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}else{
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}
