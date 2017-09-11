// findlinks1 は標準入力から読み込まれた HTML ドキュメント内の
// リンクを表示します。
package main

import (
	"os"
	"fmt"
	"golang.org/x/net/html"
)

// 実行例：
//
//   > go build ./ch1/fetch
//   > go build ./ch5/findlinks1
//   > fetch https://golang.org | findlinks1
//
func main(){
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for i, link := range visit(nil, doc){
		fmt.Printf("%2d: %s\n", i, link)  // リンク数も分かるように少し修正
	}
}

// visit は、n 内で見つかったリンクを一つひとつ links へ追加し、
// その結果を返します。
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}