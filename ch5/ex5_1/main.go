// 【練習問題 5.1】
// ループの代わりに visit への再帰呼び出しを使って n.FirstChild
// リンクリストを走査するように findlinks プログラムを変更しなさい。
package main

import (
	"os"
	"fmt"
	"golang.org/x/net/html"
)

// 実行例：
//
//   > go build ./ch1/fetch
//   > go build ./ch5/ex5_1
//   > fetch https://golang.org | ex5_1
//
func main(){
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex5_1: %v\n", err)
		os.Exit(1)
	}

	for i, link := range visit(nil, doc){
		fmt.Printf("%2d: %s\n", i, link)
	}
}

// visit は、n 内に子ノードがあればそれを走査し、そのノードを除去して
// 再度 n を走査します。　子ノードが無くなったら n 自身を処理します。
// （この実装より下記のコメントアウトしているものの方が、
// 子ノードを削除せず効率的でエレガント！）
func visit(links []string, n *html.Node) []string {
	if c := n.FirstChild; c != nil {
		// 子ノードがある場合はそれを走査して、そのノードを除去し、
		// 再度このノードを走査する
		links = visit(links, c)
		n.RemoveChild(c)
		links = visit(links, n)

	}else{
		// 子ノードが無くなったら、このノードが <a/> の場合に href 属性を links に追加する
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
	}

	return links
}

//// 下記のコードを参考にしました：
////
////   https://github.com/eliben/go-samples/blob/master/gopl-exercises/ch5-findlinks.go
////
//func visit(links []string, n *html.Node) []string {
//	if n == nil {
//		return links
//	}
//
//	if n.Type == html.ElementNode && n.Data == "a" {
//		for _, a := range n.Attr {
//			if a.Key == "href" {
//				links = append(links, a.Val)
//			}
//		}
//	}
//
//	links = visit(links, n.FirstChild)
//	return visit(links, n.NextSibling)
//}