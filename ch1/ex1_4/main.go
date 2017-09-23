// 【練習問題 1.4】
// 重複した行のそれぞれが含まれていたすべてのファイルの名前を表示するように
// dup2 を修正しなさい。
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 実行例：
//
//   > go run ./ch1/ex1_4/main.go ./resources/commitors1.txt ./resources/commitors2.txt ^
//   > ./resources/commitors3.txt ./resources/commitors4.txt
//
// PowerShell や bash のリダイレクトを使えばもっと短くできますが。
//
// commitors#.txt ファイルは、golang の GitHub から、適当な日のコミット実行者を拝借してリストアップしたものです。
// また、go get でコードを取得した場合は、上記のコマンドではうまくいかないかもしれません。
func main() {
	counts := make(map[string]int)
	containingFiles := make(map[string]string)
	// 重複した文（counts のキーと同じ）とそれを含んでいるファイル名のマップ
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		localCounts := make(map[string]int)
		// 個々のファイル内での重複数を保持する。
		for input.Scan() {
			s := input.Text()
			counts[s]++
			localCounts[s]++

			// 最初に文が現れたときのみ、containingFiles にファイル名を追加
			if localCounts[s] == 1 {
				containingFiles[s] += "[" + arg + "]"
			}
		}
		f.Close()
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s %s\n", n, line, containingFiles[line])
		}
	}
}
