// 【練習問題 5.3】
// HTML ドキュメントツリー内のすべてのテキストノードの内容を表示する
// 関数を書きなさい。ウェブブラウザでは内容が表示されない <script> と
// <style> 要素の中は調べないようにしなさい。
package main

import (
	"os"
	"fmt"
	"golang.org/x/net/html"
	"github.com/waman/exercise-go/ch5/htmlutil"
	"log"
	"io"
)

// 次節の findlinks2 を参考にして第1章の fetch を実行しなくてよいようにしています。
// 指定した URL から HTML を読み込むコードは定型文なので htmlutil.GetHTML に
// 抽出しました。
//
// 実行例：
//
//   > go run ex5_3 https://golang.org
//
func main(){
	doc, err := htmlutil.GetHTML(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "練習問題 5.3： %v\n", err)
		os.Exit(1)
	}

	visit(os.Stdout, doc)
}

// visit は要素 n 内のテキストノードを w に書き出します。
// ただし <script> と <style> に含まれるテキストノードは書き出しません。
// 子要素に対しては visit でさらに走査を行います。
func visit(w io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" { return }

	}else if n.Type == html.TextNode {
		fmt.Fprint(w, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(w, c)
	}
}