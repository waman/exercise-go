// 【練習問題 5.16】
// strings.Join の可変個引数のバージョンを書きなさい。
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Printf(join(os.Args[1:]...))
}

func join(ss ...string) string {
	var buf bytes.Buffer
	for _, s := range ss {
		buf.WriteString(s)
	}
	return buf.String()
}
