// 【練習問題 4.4】
// 一回のパスで操作を行う RotateLeft を書きなさい。
package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

// 最後の引数を除く引数を入力の文字列スライスとし、整数を与える最後の引数分だけ回転させます。
// 実行例：
//
//   > go run ./ch4/ex4_4/main.go a b c d e 3
//   > [d e a b c]
//
func main(){
	// コマンドライン引数のコピーを作成
	n := len(os.Args)
	a := make([]string, n-2)
	copy(a, os.Args[1:n-1])

	i, err := strconv.Atoi(os.Args[n-1])
	if err != nil  || i < 0 {
		log.Fatalf("最後の引数は自然数にしてください： %s %s", os.Args[n-1], err)
		return
	}

	// 回転の実行
	RotateLeft(a, i)

	fmt.Println(a)
}

func RotateLeft(s []string, n int) {
	N := len(s)
	m := n%N
	if N == 0 || N == 1 || m == 0 {
		return
	}

	rotate(s, m)
}

func RotateRight(s []string, n int) {
	N := len(s)
	m := n%N
	if N == 0 || N == 1 || m == 0 {
		return
	}

	rotate(s, N-m)
}

func rotate(s []string, n int){
	N := len(s)
	i := 0
	for m := N-n; i < m; i++ {  // i+n = N-1 まで
		s[i], s[i+n] = s[i+n], s[i]
	}

	if m := N%n; m != 0 {
    rotate(s[i:], n-m)
	}
}