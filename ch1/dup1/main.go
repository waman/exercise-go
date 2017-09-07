// dup1 は標準入力から2回以上現れる行を出現回数と一緒に表示します。
package main

import (
	"bufio"
	"os"
	"fmt"
)

// 実行例：
//
//   > cd ch1
//   > go run ./dup1/main.go < commitors1.txt
//
// commitors#.txt ファイルは、golang の GitHub から、適当な日のコミット実行者を拝借してリストアップしたものです。
// また、go get でコードを取得した場合は、上記のコマンドではうまくいかないかもしれません。
func main(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
