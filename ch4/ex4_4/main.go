// 【練習問題 4.4】
// 一回のパスで操作を行う rotate を書きなさい。
package main

import (
	"fmt"
	"os"
)

func main(){
	// コマンドライン引数のコピーを作成
	a := make([]string, len(os.Args)-1)
	copy(a, os.Args[1:])

	// 回転の実行
	rotate(a)

	fmt.Println(a)
}

func rotate(s []string) {
	if n := len(s); n == 0 || n == 1 { return }

	head, n := s[0], len(s)-1
	for i := 0; i < n; i++ {
		s[i] = s[i+1]
	}
	s[n] = head
}
