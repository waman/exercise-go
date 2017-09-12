package main

import (
	"golang.org/x/net/html"
	"fmt"
	"github.com/waman/exercise-go/ch5/links"
	"net/http"
	"strings"
	"os"
	"log"
)

func main(){
	title, err := title(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(title)
}

func title(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()  // defer!

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;"){
		return "", fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return soleTitle(doc)
}

// soleTitle は doc の中の最初の空でない title 要素のテキストと、
// title 要素が一つだけでなかったらエラーを返します。
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func(){
		switch p := recover(); p {
		case nil:
			// パニックなし
		case bailout{}:
			// 「予期された」パニック
			err = fmt.Errorf("multiple title elements")
		default:
			// 予期しないパニック
			 panic(p)
		}
	}()

	links.ForEachNode(doc, func(n *html.Node){
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
				if title != "" {
					panic(bailout{})
				}
				title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("notitle element")
	}
	return title, nil
}