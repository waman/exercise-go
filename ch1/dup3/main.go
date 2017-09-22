// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 実行例：
//
//   > cd ch1
//   > go run ./dup2/main.go commitors1.txt commitors2.txt
//
// commitors#.txt ファイルは、golang の GitHub から、適当な日のコミット実行者を拝借してリストアップしたものです。
// また、go get でコードを取得した場合は、上記のコマンドではうまくいかないかもしれません。
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
