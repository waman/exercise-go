// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// forEachNode は n から始まるツリー内の個々のノード x に対して
// 関数 pre(x) と post(x) を呼び出します。その二つの関数はオプションです。
// pre は子ノードを訪れる前に呼び出され（前順 : preorder）、
// post は子ノードを訪れた後に呼び出されます（降順 : postorder）。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

// 実行例：
//
//   > go build ./ch1/fetch
//   > go build ./ch5/outline2
//   > fetch https://golang.org | outline
//
func main() {
	url := os.Args[1]

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("parsing %s: as HTML: %v", url, err)
	}
	forEachNode(doc, startElement, endElement)
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
