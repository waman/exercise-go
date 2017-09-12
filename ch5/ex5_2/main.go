// 【練習問題 5.2】
// p、div、span などの要素名に、HTML ドキュメントツリー内でその要素名を持つ
// 要素の数を対応させるマッピングを行う関数を書きなさい。
package main

import (
	"os"
	"fmt"
	"golang.org/x/net/html"
)

// 実行例：
//
//   > go build ./ch1/fetch
//   > go build ./ch5/ex5_2
//   > fetch https://golang.org | ex5_2
//
func main(){
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "練習問題 5.2： %v\n", err)
		os.Exit(1)
	}

	for tag, count := range visit(nil, doc){
		fmt.Printf("%2d: %s\n", tag, count)
	}
}

// visit は、n 内で見つかったリンクを一つひとつ links へ追加し、
// その結果を返します。
func visit(tagCounts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		tagCounts[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tagCounts = visit(tagCounts, c)
	}

	return tagCounts
}