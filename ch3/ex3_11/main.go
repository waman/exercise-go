// 【練習問題 3.11】
// comma を機能拡張して、符号記号を持つ浮動小数点数を
// 正しく扱えるようにしなさい。
package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

// コマンドライン引数（なければ標準入力）で与えられた数値に comma を
// 適用して表示する
func main(){
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
	// 符号
	sign := ""
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+"){
		sign = s[0:1]
		s = s[1:]
	}

	// 整数部最後のインデックス
	i := strings.IndexAny(s, ".eE")

	if i == -1 {
		return sign + insertComma(s)
	}else {
		return sign + insertComma(s[:i]) + s[i:]
	}
}

// 本文の comma と同じ
func insertComma(s string) string {
	n := len(s)
	if n <= 3 { return s }
	return insertComma(s[:n-3]) + "," + s[n-3:]
}
