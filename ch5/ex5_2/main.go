// 【練習問題 5.2】
// p、div、span などの要素名に、HTML ドキュメントツリー内でその要素名を持つ
// 要素の数を対応させるマッピングを行う関数を書きなさい。
package main

import (
	"os"
	"fmt"
	"golang.org/x/net/html"
	"github.com/waman/exercise-go/ch5/htmlutil"
	"log"
)

// 次節の findlinks2 を参考にして第1章の fetch を実行しなくてよいようにしています。
// 指定した URL から HTML を読み込むコードは定型文なので htmlutil.GetHTML に
// 抽出しました。
//
// 実行例：
//
//   > go run ex5_2 https://golang.org
//
func main(){
	doc, err := htmlutil.GetHTML(os.Args[1])
	if err != nil {
		log.Fatalf("練習問題 5.2： %v\n", err)
	}

	var tagCounts = make(map[string]int)
	visit(tagCounts, doc)

	for tag, count := range tagCounts{
		fmt.Printf("%s:\t %d\n", tag, count)
	}
}

// visit は要素の名前を tagCounts に追加し、子要素に対して visit を呼び出します。
func visit(tagCounts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		tagCounts[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(tagCounts, c)
	}
}