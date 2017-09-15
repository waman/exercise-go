// 【練習問題 5.17】
// HTML ノードツリーと0個以上の名前が与えられたら、それらの名前の一つと位置する
// 要素をすべて返す可変個引数関数 ElementsByTagName を書きなさい。二つの呼び出し
// 例を次に示します。
//
//   func ElementByTagName(doc *html.Node, name ...string) []*html.Node
//
//   images := ElementsByTagName(doc, "img")
//   headings := ElementByTagName(doc, "h1", "h2", "h3", "h4")
//
package main

import (
	"golang.org/x/net/html"
	"log"
	"fmt"
	"github.com/waman/exercise-go/ch5/htmlutil"
	"os"
)

func main(){
	url := os.Args[1]
	tags := os.Args[2:]

	doc, err := htmlutil.GetHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	nodes := ElementByTagName(doc, tags...)
	fmt.Printf("%d 個の要素が見つかりました。\n", len(nodes))
}

func ElementByTagName(doc *html.Node, tags ...string) []*html.Node {
	var result []*html.Node

	pre := func(n *html.Node){
		if n.Type == html.ElementNode {
			eName := n.Data
			for _, tag := range tags {
				if eName == tag {
					result = append(result, n)
				}
			}
		}
	}

	htmlutil.ForEachNode(doc, pre, nil)

	return result
}