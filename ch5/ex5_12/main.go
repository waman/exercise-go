// 【練習問題 5.12】
// gopl.io/ch5/outline2 （5.5節）の startElement 関数と endElement 関数は
// グローバル変数 depth を共有しています。その二つの関数を無名関数にして、
// outline に対してローカルな変数を共有するようにしなさい。
package main

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
	"log"
	"github.com/waman/exercise-go/ch5/htmlutil"
)

// findlinks2 を参考にして第1章の fetch を実行しなくてよいようにしています。
// 指定した URL から HTML を読み込むコードは定型文なので htmlutil.GetHTML に
// 抽出しました。
//
// 実行例：
//
//   > go run ex5_4 https://golang.org
//
func main(){
	doc, err := htmlutil.GetHTML(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	outline(doc)
}

func outline(doc *html.Node){
	var depth int

	startElement := func(n *html.Node){
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	endElement := func(n *html.Node){
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	// htmlutil.ForEachNode 関数は outline2 の forEachNode 関数と同じです。
	// 他の練習問題でも使うので公開された関数として定義しています。
	htmlutil.ForEachNode(doc, startElement, endElement)
}
