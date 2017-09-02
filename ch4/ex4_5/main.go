// 【練習問題 4.5】
// []string スライス内で隣接している重複をスライス内で除去する
// 関数を書きなさい。
package main

import (
	"os"
	"fmt"
)

func main(){
	// コマンドライン引数のコピーを作成
	a := make([]string, len(os.Args)-1)
	copy(a, os.Args[1:])

	a = removeContinuousDuplicates(a)
	fmt.Println(a)
}

func removeContinuousDuplicates(s []string) []string{
	n := len(s)
	if n == 0 { return s }

	current, nextIndex := s[0], 1
	for i := 1; i < n; i++ {
		if s[i] != current {
			current = s[i]
			s[nextIndex] = current
			nextIndex++
		}
	}

	return s[:nextIndex]
}