package htmlutil

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

// GetHTML は url に対して HTTP GET リクエストを行い、レスポンスを
// HTML としてパースして *html.Node を返します。
func GetHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 5.8節「遅延関数呼び出し」参照

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s の取得に失敗しました： %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s を HTML としてパースするのに失敗しました： %v", url, err)
	}

	return doc, nil
}

// outline2 のものと同じ。　ただしパブリックにしています。
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

// BreadthFirst は findLinks3 の breadthFirst と同じです。
func BreadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
