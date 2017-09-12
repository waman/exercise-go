// パッケージ links はリンク抽出関数を適用します。
package links

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
)

// Extract は指定された URL へ HTTP GET リクエストを行い、
// レスポンスを HTML としてパースして、その HTML ドキュメント
// 内のリンクを返します。
func Extract(url string)([]string, error) {
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
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node){
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue  // 無視
				}
				links = append(links, link.String())
			}
		}
	}
	ForEachNode(doc, visitNode, nil)
	return links, nil
}

// outline2 のものと同じ。　ただしパブリックにしています。
func ForEachNode(n *html.Node, pre, post func(n *html.Node)){
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
