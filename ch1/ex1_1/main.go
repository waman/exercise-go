// 【練習問題 1.1】
// echo プログラムを修正して、そのプログラムを起動したコマンド名である
// os.Args[0] も表示するようにしなさい。
package main

import (
	"os"
	"fmt"
)

func main(){
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}