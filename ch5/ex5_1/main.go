// 【練習問題 5.1】
// ループの代わりに visit への再帰呼び出しを使って n.FirstChild
// リンクリストを走査するように findlinks プログラムを変更しなさい。
package main

import (
	"os"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"github.com/waman/exercise-go/ch5/htmlutil"
)

// 次節の findlinks2 を参考にして第1章の fetch を実行しなくてよいようにしています。
// 指定した URL から HTML を読み込むコードは定型文なので htmlutil.GetHTML に
// 抽出しました。
//
// 実行例：
//
//   > go run ex5_1 https://golang.org
//
func main(){
	doc, err := htmlutil.GetHTML(os.Args[1])
	if err != nil {
		log.Fatalf("ex5_1: %v\n", err)
	}

	for i, link := range visit(nil, doc){
		fmt.Printf("%2d: %s\n", i, link)
	}
}

// 下記のコードを参考にしました：
//
//   https://github.com/eliben/go-samples/blob/master/gopl-exercises/ch5-findlinks.go
//
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visit(links, n.FirstChild)
	return visit(links, n.NextSibling)
}

//// 破壊的走査の方法www
//func visit(links []string, n *html.Node) []string {
//	if c := n.FirstChild; c != nil {
//		// 子ノードがある場合はそれを走査して、そのノードを除去し、
//		// 再度このノードを走査する
//		links = visit(links, c)
//		n.RemoveChild(c)
//		links = visit(links, n)
//
//	}else{
//		// 子ノードが無くなったら、このノードが <a/> の場合に href 属性を links に追加する
//		if n.Type == html.ElementNode && n.Data == "a" {
//			for _, a := range n.Attr {
//				if a.Key == "href" {
//					links = append(links, a.Val)
//				}
//			}
//		}
//	}
//
//	return links
//}