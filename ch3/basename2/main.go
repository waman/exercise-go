// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// コマンドライン引数（なければ標準入力）で与えられた数値に basename を
// 適用して表示する
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)

		fmt.Print("> ")
		for input.Scan() {
			fmt.Println(basename(input.Text()))
			fmt.Print("> ")
		}
	} else {
		for _, arg := range args {
			fmt.Println(basename(arg))
		}
	}
}

// basename はディレクトリ要素と . 接尾辞を取り除きます。
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
