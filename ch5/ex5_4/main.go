// 【練習問題 5.4】
// visit 関数を拡張して、画像、スクリプト、スタイルシートなどの
// 他の種類のリンクをドキュメントから抽出するようにしなさい。
package main

import (
	"fmt"
	"github.com/waman/exercise-go/ch5/htmlutil"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
)

// 次節の findlinks2 を参考にして第1章の fetch を実行しなくてよいようにしています。
// 指定した URL から HTML を読み込むコードは定型文なので htmlutil.GetHTML に
// 抽出しました。
//
// 実行例：
//
//   > go run ex5_4 https://golang.org
//
func main() {
	doc, err := htmlutil.GetHTML(os.Args[1])
	if err != nil {
		log.Fatalf("練習問題 5.4： %v\n", err)
	}

	visit(os.Stdout, doc)
}

var linkAtts = map[string]string{
	"img":    "src",
	"script": "src",
	"link":   "href",

	"table": "background",
	"td":    "background",
	"th":    "background",
	"tr":    "background",

	"input": "src",

	"object": "data",
	"audio":  "src",
	"track":  "src",
	"video":  "src", // "video":"poster"
	"source": "src",
}

// visit は要素 n 内のリンクを w へ書き出します。
// 子要素に対しては visit でさらに走査を行います。
func visit(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		att := linkAtts[n.Data]
		if attVal := getAttribute(n, att); attVal != "" {
			fmt.Fprintf(w, "%s\n", attVal)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(w, c)
	}
}

// 属性がない場合には、複数返り値で bool 値の ok を返す方がいい気がしますが、
// まだ出てきてないので空文字列を返します。
func getAttribute(n *html.Node, name string) string {
	for _, a := range n.Attr {
		if a.Key == name {
			return a.Val
		}
	}
	return ""
}
