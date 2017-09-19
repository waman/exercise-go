// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"os"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

// 実行例：
//
//   > go run ./ch5/findlinks2/main.go https://golang.org https://github.com/golang
//
// パースしたいサイトの URL を複数指定できます。
// また、findlinks1 のように ch1/fetch を使う必要はありません。
func main(){
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks は url に対して HTTP GET リクエストを行い、レスポンスを
// HTML としてパースして、リンクを抽出して返します。
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s: as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// findlinks1 の visit と同じ。
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