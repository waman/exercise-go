// 【練習問題 7.2】
// 下記のシグニチャを持つ関数 CountingWriter を書きなさい。io.Writer が
// 与えられたなら、それを包む新たな Writer と int64 変数へのポインタを返します。
// その変数は新たな Writer に書き込まれたバイト数を常に保持しています。
//
//   func CountingWriter(w io.Writer) (io.Writer, *int64)
//
package main

import (
	"io"
	"bytes"
	"fmt"
)

type countingWriter struct {
	io.Writer
	n *int64
}

func (cw *countingWriter) Write(p []byte) (n int, err error) {
	*cw.n += int64(len(p))
	return cw.Writer.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var n int64 = 0
	cw := &countingWriter{w, &n}
	return cw, &n
}

func main(){
	var buf bytes.Buffer
	w, n := CountingWriter(&buf)

	w.Write([]byte("Hello"))
	fmt.Println(buf.String())   // "Hello"
	fmt.Println(*n)             // 5

	w.Write([]byte(", World!"))
	fmt.Println(buf.String())   // "Hello, World!"
	fmt.Println(*n)             // 13
}