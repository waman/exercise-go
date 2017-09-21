// 【練習問題 5.13】
// crawl を修正して、必要に応じてディレクトリを作成しながら見つけたページの
// 複製をローカルに作成するようにしなさい。異なるドメインのページの複製は
// しないようにしなさい。たとえば、もとのページが golang.org からであれば
// そこにあるすべてのファイルは保存しますが、vimeo.com からのファイルは
// 保存しないということです。
package main

import (
	"bytes"
	"fmt"
	"github.com/waman/exercise-go/ch5/htmlutil"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	. "net/url" // 接頭辞なしで使えるようにする
	"os"
	"path"
	"strings"
)

var host string // host = "abc"

// あまり完成していませんがこのくらいで。
func main() {
	if len(os.Args) != 2 {
		log.Fatal("指定できる URL は1つです。")
	}

	startURL := os.Args[1]

	url, err := Parse(startURL)
	if err != nil {
		log.Fatal(err)
	}
	host = url.Host

	log.Printf("%s からリンクされている %s 内のファイルをクロールします。",
		url, url.Scheme+"://"+host)

	htmlutil.BreadthFirst(crawl, []string{startURL})
}

// net/url パッケージの練習も兼ねて *URL 型を使っています。
func crawl(urlStr string) []string {
	url, err := Parse(urlStr) // url は *URL 型
	if err != nil {
		return nil
	}

	// ドメインが等しくなければ return
	if url.Host != host {
		return nil
	}

	url, contentType, content, err := readAll(url)
	if err != nil {
		log.Print(err)
		return nil
	}

	if err = output(url, content); err != nil {
		log.Print(err)
		return nil
	}

	if strings.Contains(contentType, "text/html") {
		links, err := extractLinks(url, content)
		if err != nil {
			log.Print(err)
		}
		return links

	} else {
		return nil
	}
}

// 第1返り値の *URL は、リダイレクトなどで URL が変更された場合に使います。
// 第2返り値は Content-Type です。
func readAll(url *URL) (*URL, string, []byte, error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return nil, "", nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", nil, fmt.Errorf("%s の取得に失敗しました： %s", url, resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	return resp.Request.URL, resp.Header.Get("Content-Type"), content, err
}

func output(url *URL, content []byte) error {
	file := url.Host + url.Path // file = "abc/def/ghi.html"
	if strings.HasSuffix(file, "/") {
		file += "index.html"
	}

	dir := path.Dir(file) // dir = "abc/def/"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	fmt.Printf("%s => %s\n", url, file)

	return nil
}

func extractLinks(url *URL, content []byte) ([]string, error) {
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode { // <a> 要素に限らず、@href, @src があれば取得
			for _, a := range n.Attr {
				if a.Key == "href" || a.Key == "src" {
					link, err := url.Parse(a.Val)
					if err != nil {
						continue
					}
					links = append(links, link.String())
				}
			}
		}
	}

	htmlutil.ForEachNode(doc, visitNode, nil)
	if err != nil {
		return links, err
	}
	return links, nil
}
