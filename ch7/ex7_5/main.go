// 【練習問題 7.5】
// io パッケージの LimitReader 関数は io.Reader である r とバイト数 n を
// 受け取り、r から読み出す別の Reader を返しますが、n バイトを読み出した後に
// ファイルの終わりの状態を報告します。その関数を実装しなさい。
//
//   func LimitReader(r io.Reader, n int64) io.Reader
//
package main

import (
	"io"
	"os"
	"strings"
)

type limitReader struct {
	io.Reader
	rest int64
}

func (lr *limitReader) Read(p []byte) (int, error){
	if np := int64(len(p)); np <= lr.rest {
		lr.rest -= np
		return lr.Reader.Read(p)

	}else{
		n, err := lr.Read(p[:lr.rest])
		lr.rest = 0
		if err != nil {
			return n, err
		}
		return n, io.EOF
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}

func main(){
	sr := strings.NewReader("abcdefghijklmnopqrstuvwxyz")
	r := LimitReader(sr, 9)
	io.Copy(os.Stdout, r)
}