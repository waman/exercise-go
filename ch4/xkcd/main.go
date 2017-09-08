// ./ch4/ex4_12/main.go によって作られたインデックスファイルを使って、
// コマンドラインに渡されたキーワードをタイトルに含むコミックの URL と
// 内容 (transcript) を表示します。
package main

import (
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
	"strings"
	"os"
)

type Comic struct {
	Number int    `json:"num"`
	Title  string `json:"title"`
	Alt    string `json:"alt"`
	Image  string `json:"img"`
}

// 使用例：
//
//   > go run ./ch4/xkcd/main.go Super
//
func main() {
	indexFile := "xkcd-index.json"

	index, err := readIndex(indexFile)
	if err != nil { log.Fatalf("インデックスファイルの読み込みに失敗しました(%s): %s", indexFile, err) }

	keyword := os.Args[1]
	for _, comic := range *index {
		if strings.Contains(comic.Title, keyword) {
			fmt.Printf("***** %s [%s] *****\n", comic.Title, comic.Image)
			fmt.Printf("%s\n", comic.Alt)
			fmt.Println()
		}
	}
}

func readIndex(s string) (*[]Comic, error) {
	bs, err := ioutil.ReadFile(s)
	if err != nil { return nil, err }

	var index []Comic
	err = json.Unmarshal(bs, &index)
	if err != nil { return nil, err }

	return &index, nil
}