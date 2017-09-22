// 【練習問題 1.2】
// echo プログラムを修正して、個々の引数のインデックスと値の組を
// 1行ごとに表示しなさい。
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, ":", arg)
		// Println() にはコンマで区切って任意個の引数を渡せる
	}
}
