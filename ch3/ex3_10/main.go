// 【練習問題 3.10】
// 文字列の連結の代わりに bytes.Buffer を使って、
// 再帰呼び出しを行わない comma を作成しなさい。
package main

import (
	"bufio"
	"bytes"
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

func comma(s string) string {
	length := len(s)
	if length < 3 {
		return s
	}

	var buf bytes.Buffer

	i := length % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])
	s = s[i:]

	for ; len(s) > 0; s = s[3:] {
		buf.WriteByte(',')
		buf.WriteString(s[:3])
	}

	return buf.String()
}
