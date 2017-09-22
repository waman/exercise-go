// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"bufio"
	"fmt"
	"os"
)

// コマンドライン引数（なければ標準入力）で与えられた数値に comma を
// 適用して表示する
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)

		fmt.Print("> ")
		for input.Scan() {
			fmt.Println(comma(input.Text()))
			fmt.Print("> ")
		}
	} else {
		for _, arg := range args {
			fmt.Println(comma(arg))
		}
	}
}

// comma は負ではない10進表記整数文字列にカンマを挿入します。
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
