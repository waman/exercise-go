package main

import (
	"testing"
	"github.com/waman/exercise-go/ch5/htmlutil"
	"strings"
	"bytes"
	"golang.org/x/net/html"
)

func Testプリティプリンタで出力したHTMLがパースできる(t *testing.T){
	//url := "https://golang.org"
	url := "http://gopl.io"
	doc, _ := htmlutil.GetHTML(url)

	var sw bytes.Buffer
	PrettyPrint(doc, &sw)
	output := sw.String()

	_, err := html.Parse(strings.NewReader(output))
	// HTML として壊れている output を渡してもパースが通ってしまうのだが・・・
	if err != nil {
		t.Errorf("出力したドキュメントのパースに失敗しました：%s %v", url, err)
	}
}
