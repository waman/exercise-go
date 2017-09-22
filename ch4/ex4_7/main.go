// 【練習問題 4.7】
// UTF-8 でエンコードされた文字列を表す []byte スライスの文字を、
// そのスライス内で逆順にするように reverse を修正しなさい。
// 新たなメモリを割り当てることなく行えるでしょうか。
package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	// 引数をバイトスライスへ変換
	bs := ToByteSlice(os.Args[1:])

	// 反転の実行
	reverse(bs)

	fmt.Println(bs)
}

// 文字列のスライスをバイトスライスへ変換
func ToByteSlice(strslice []string) []byte {
	var buffer bytes.Buffer
	for _, arg := range strslice {
		buffer.WriteString(arg)
	}
	return buffer.Bytes()
}

func reverse(bs []byte) {
	n := len(bs)
	if n == 0 || n == 1 {
		return
	}

	for i := 0; i < n; {
		_, size := utf8.DecodeRune(bs[i:])
		reverseBytes(bs[i : i+size]) // ルーン内の反転
		i += size
	}
	reverseBytes(bs) // 全体の反転
}

// []byte スライスを byte 列として反転する。
// 本文の reverse の []byte 版。
func reverseBytes(bs []byte) {
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
}
