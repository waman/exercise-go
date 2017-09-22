// 【練習問題 4.6】
// UTF-8 でエンコードされた []byte スライス内で隣接している
// Unicode スペース（unicode.IsSpace を参照）を、もとの
// スライス内で一つの ASCII スペースへ圧縮する関数を書きなさい。
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	// 引数をバイトスライスへ変換
	bs := ToUnquotedByteSlice(os.Args[1:])

	// 空白の圧縮
	bs = compressWhiteSpaces(bs)

	fmt.Println(bs)
}

// 引数が [Hello, Go \u4e16\u754c!] のとき
// 以下の unquoted （の列）は [Hello, Go 世界!] となります。
// つまり、空白文字を含めて、ユニコードエスケープした文字列で文字を指定できます。
//
// bytes は [72 101 108 108 111 44 71 111 228 184 150 231 149 140 33]
//           H  e   l   l   o   ,  G  o   世          界           !
func ToUnquotedByteSlice(strslice []string) []byte {
	var buffer bytes.Buffer
	for _, arg := range strslice {
		unquoted, err := strconv.Unquote("\"" + arg + "\"")
		if err != nil {
			fmt.Println(err)
		}

		buffer.WriteString(unquoted)
	}
	return buffer.Bytes()
}

func compressWhiteSpaces(bs []byte) []byte {
	n := len(bs)
	if n == 0 {
		return bs
	}

	previousIsSpace, nextIndex := false, 0
	for i := 0; i < n; {
		r, size := utf8.DecodeRune(bs[i:]) // 位置iからルーンを1文字読み込む
		if unicode.IsSpace(r) {
			if !previousIsSpace {
				bs[nextIndex] = ' '
				nextIndex++
			}
			previousIsSpace = true

		} else {
			copy(bs[nextIndex:], bs[i:i+size]) // 読み込んだルーンのバイト数分を前へコピー
			nextIndex += size
			previousIsSpace = false
		}
		i += size
	}

	return bs[:nextIndex]
}
