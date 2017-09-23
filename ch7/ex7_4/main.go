// 【練習問題 7.4】
// strings.NewReader 関数は、その引数である文字列から読み込むことで
// io.Reader インタフェース（と他のインタフェース）を満足する値を返し
// ます。みなさん自身で簡単な NewReader を実装し、HTML パーサ（5.2節）
// が文字列からの入力を受け取るようにしなさい。
package main

import (
	"io"
	"golang.org/x/net/html"
	"fmt"
	"log"
	"github.com/waman/exercise-go/ch5/htmlutil"
	"os"
)

type stringReader struct {
	content []byte
}

func (sr *stringReader) Read(p []byte) (int, error) {
	if n, np := len(sr.content), len(p); n >= np {
		copy(p, sr.content[:np])
		sr.content = sr.content[np:]
		return np, nil

	}else{
		copy(p[:n], sr.content)
		sr.content = nil
		return n, io.EOF
	}
}

func NewReader(s string) *io.Reader {
	var sr io.Reader = &stringReader{[]byte(s)}
	return &sr
}

func main(){
	// 文字列を読み込む
	content := "abcdef"
	r := NewReader(content)
	io.Copy(os.Stdout, *r)
	fmt.Println()

	// HTML を読み込む
	content = "<html><body><a href='https://golang.org'>Google Go</a></body></html>"
	r = NewReader(content)

	doc, err := html.Parse(*r)
	if err != nil {
		log.Fatal(err)
	}

	pre := func(n *html.Node){
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, att := range n.Attr {
				if att.Key == "href" {
					fmt.Println(att.Val)
				}
			}
		}
	}
	htmlutil.ForEachNode(doc, pre, nil)
}
