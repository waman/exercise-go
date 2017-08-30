// 【練習問題 3.10】
// 文字列の連結の代わりに bytes.Buffer を使って、
// 再帰呼び出しを行わない comma を作成しなさい。
package main

import (
	"os"
	"bufio"
	"fmt"
	"bytes"
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

func comma(s string) string {
	ss := []byte(s)
	var buf bytes.Buffer

	i := len(ss)%3
	buf.Write(ss[:i])
	ss = ss[i:]

	for ; len(ss) > 0; ss = ss[3:] {
		buf.WriteByte(',')
		buf.Write(ss[:3])
	}

	return buf.String()
}

