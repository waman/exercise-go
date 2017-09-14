// 【練習問題 5.9】
// 文字列 s 内のそれぞれの部分文字列 "$foo" を f("foo") が返すテキスト*（訳注↓）で
// 置換する関数 expand(s string, f func(string)string) string を書きなさい。
// （訳注：$ で始まる任意の単語を探して、$ 以降の文字列で関数 f を呼び出した結果のテキストです。）
package main

import (
	"strings"
	"fmt"
	"bytes"
	"unicode"
)

func main(){
	upper := func(s string) string {
		return strings.ToUpper(s)
	}
	fmt.Println(expand("Hello, $world!", upper))  // "Hello, WORLD!"
}

func expand(s string, f func(string)string) string {
  var buffer bytes.Buffer
  expandToWriter(&buffer, s, f)
  return buffer.String()
}

// 本文で出てきた Scanner を単語で分割する方法では、空白が維持できないので
// 用いていません。
func expandToWriter(buf *bytes.Buffer, s string, f func(string)string){
  i := strings.Index(s, "$")
  if i == -1 {
  	buf.WriteString(s)
  	return
	}

	// s = "Hello, $world!"
	buf.WriteString(s[:i]) // "Hello, " を書き出す
	s = s[i+1:]            // s = "world!"

	j := strings.IndexFunc(s, IsNotLetter )
	word := s[:j]    // word = "world"
	buf.WriteString(f(word))

	s = s[j:]        // s = "!"
	expandToWriter(buf, s, f)
}

func IsNotLetter(r rune) bool {
	return !unicode.IsLetter(r)
}