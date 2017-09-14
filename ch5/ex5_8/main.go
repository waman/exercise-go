// 【練習問題 5.8】
// 走査を続けるか否かを示すブーリアンの結果を pre 関数と post 関数が返す
// ようにして、それに対応するように forEachNode を修正しなさい。修正した
// forEachNode を使って、指定された id 属性を持つ最初の HTML 要素を見つける
// ような下記のシグニチャの関数 ElementByID を書きなさい。ElementByID は、
// 一致が見つかったら走査を中止しなければなりません。
//
//   func ElementByID(doc *html.Node, id string) *html.Node
//
package main

import (
	"golang.org/x/net/html"
	"strings"
	"log"
	"fmt"
)

func main(){
	text := `<ol><li id="1">first</li><li id="2">second</li><li id="3">third</li></ol>`
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		log.Fatal(err)
	}

	node := ElementByID(doc, "2")
	fmt.Println(node.FirstChild.Data)  // second
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var target *html.Node

	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			if value, ok := getAttribute(n, "id"); ok && value == id {
				target = n
				return false
			}
		}
		return true
	}

	forEachNode(doc, pre, nil)

	return target
}

func getAttribute(n *html.Node, name string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == name {
      return a.Val, true
		}
	}
	return "", false
}

func forEachNode(n *html.Node, pre, post func(*html.Node)bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil {
		if !post(n) {
			return false
		}
	}

	return true
}
