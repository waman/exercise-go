// 【練習問題 5.7】
// 汎用の HTML プリティプリンタとなるような startElement と endElement を
// 開発しなさい。コメントノード、テキストノード、個々の要素の属性（<a href='...'>）
// を表示しなさい。要素が子を持たない場合には、<img></img> ではなく <img/> のような
// 短い形式を使いなさい。出力をきちんとパースできることを保障するためのテストを書きなさい
// （11章参照）。
package main

import (
	"github.com/waman/exercise-go/ch5/htmlutil"
	"os"
	"log"
	"golang.org/x/net/html"
	"io"
	"fmt"
)

// 実行例：
//
//   > go run ./ch5/ex5_7/main.go http://gopl.io
//
// テスト実行例：
//
//    > go test ./ch5/ex5_7
//
func main(){
	doc, err := htmlutil.GetHTML(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

  PrettyPrint(doc, os.Stdout)
}

func PrettyPrint(n *html.Node, w io.Writer){
	pre := func(n *html.Node){
		switch n.Type {
		case html.DoctypeNode:
			fmt.Fprintf(w, "<!DOCTYPE %s>", n.Data)

		case html.ElementNode:
			fmt.Fprintf(w, "<%s", n.Data)

			for _, a := range n.Attr {
				fmt.Fprintf(w, " %s='%s'", a.Key, a.Val)
			}

			// 子要素がない場合はタグを閉じる
			if n.FirstChild == nil {
				fmt.Fprint(w, "/")
			}

			fmt.Fprint(w, ">")

		case html.TextNode:
			fmt.Fprint(w, n.Data)

		case html.CommentNode:
			fmt.Fprintf(w, "<!--%s-->", n.Data)

		default:  // case html.ErrorNode:
			fmt.Errorf("Unknow node type appears: %s", n)
		}
	}

	post := func(n *html.Node){
		if n.Type == html.ElementNode && n.FirstChild != nil {
			fmt.Fprintf(w, "</%s>", n.Data)
		}
	}

  htmlutil.ForEachNode(n, pre, post)
}