// 【練習問題 4.6】
// UTF-8 でエンコードされた []byte スライス内で隣接している
// Unicode スペース（unicode.IsSpace を参照）を、もとの
// スライス内で一つの ASCII スペースへ圧縮する関数を書きなさい。
package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func compressWhiteSpaces(bytes []byte) []byte {
	n := len(bytes)
	if n == 0 { return bytes }

	previousIsSpace, nextIndex := false, 0
	for i := 0; i < n; {
		r, size := utf8.DecodeRune(bytes[i:])
		if unicode.IsSpace(r) {
			if !previousIsSpace {
				bytes[nextIndex] = ' '
				nextIndex++
			}
			previousIsSpace = true

		} else {
			copy(bytes[nextIndex:], bytes[i:i+size])
			nextIndex += size
			previousIsSpace = false
		}
		i += size
	}

	return bytes[:nextIndex]
}

func main(){
	// コマンドライン引数が [Hello, Go \u4e16\u754c!] のとき
	// 以下の unquoted （の列）は [Hello, Go 世界!] となる。
	// bytes は [72 101 108 108 111 44 71 111 228 184 150 231 149 140 33]
	//           H  e   l   l   o   ,  G  o   世          界           !
	var bytes []byte
	for _, arg := range os.Args[1:] {
		unquoted, err := strconv.Unquote("\"" + arg + "\"")
		if err != nil { fmt.Println(err) }

		for _, b := range []byte(unquoted){
			bytes = append(bytes, b)
		}
	}

	bytes = compressWhiteSpaces(bytes)
	fmt.Println(bytes)
}
