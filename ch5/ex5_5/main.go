// 【練習問題 5.5】
// countWordsAndImages を実装しなさい
// （単語の分割については練習問題 4.9 を参照）。
package main

import (
	"github.com/waman/exercise-go/ch5/htmlutil"
	"golang.org/x/net/html"
	"os"
	"fmt"
	"strings"
	"bufio"
)

func main(){
  words, images, err := CountWordsAndImages(os.Args[1])
  if err != nil {
  	fmt.Fprintf(os.Stderr, "練習問題 5.5： %v\n", err)
  	os.Exit(1)
	}

	fmt.Printf("単語数： %d\n", words)
	fmt.Printf("画像数： %d\n", images)
}

// CountWordsAndImages は HTML ドキュメントに対する HTTP GET
// リクエストを url へ行い、そのドキュメント内に含まれる単語と画像
// の数を返します。
func CountWordsAndImages(url string) (words, images int, err error) {
	doc, err := htmlutil.GetHTML(url)
	if err != nil {
		return words, images, err
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {

	var elementName string  // n が要素なら要素名、それ以外なら nil
	if n.Type == html.ElementNode {
		elementName = n.Data
	}

  if elementName == "img" {
  	images++

	}else if n.Type == html.TextNode {
    input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			words++
		}
	}

	// 練習問題 5.3 と同じく、 <script> 要素と <style> 要素の中は調べないことにします。
	if elementName != "script" && elementName != "style" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			w, i := countWordsAndImages(c)
			words += w
			images += i
		}
	}

	return
}

