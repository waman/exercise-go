// 【練習問題 4.12】
// 人気があるウェブコミック xkcd は JSON インターフェースを持っています。
// たとえば、https://xkcd.com/571/info.0.json に対するリクエストは、
// 多くのお気に入りのうちの一つであるコミック 571 の詳細な説明を生成します。
// それぞれの URL を（一度だけ！）ダウンロードして、オフラインインデックス
// を作成しなさい。そのインデックスを使って、コマンドラインで提供された検索
// 語と一致するコミックのそれぞれの URL と内容 (transcript) を表示する
// ツール xkcd を書きなさい。
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const min = 1800

var indexFile = "xkcd-index.json"

type Comic struct {
	Number int    `json:"num"`
	Title  string `json:"title"`
	Alt    string `json:"alt"` // 新しいものは transcript が空なので alt を表示
	Image  string `json:"img"`
}

// 全てのコミックを書き出すのは大変なので、最新 (No.1886) から No.1800 までを書き出してます。
//
// 実行例：
//
//   > go run ./ch4/ex4_12/main.go
//
// 作成したインデックスファイルを使用して URL と内容を表示するプログラムは xkcd 参照。
func main() {
	// 引数があればファイルに出力（os パッケージのドキュメント参照）
	file, err := os.OpenFile(indexFile, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if cErr := file.Close(); err == nil && cErr != nil {
			log.Fatal(err)
		}
	}()

	createIndex(file)
}

func createIndex(w io.Writer) {
	fmt.Fprint(w, "[")
	var sep = ""

	// current comic
	current, err := addIndexOfAComic(w, 0)
	if err != nil {
		log.Fatalf("コミック情報のインデクスに失敗しました (current): %s", err)
	} else {
		sep = ","
	}

	// older comic
	for no := current.Number - 1; no > min; no-- {
		fmt.Fprint(w, sep)
		_, err := addIndexOfAComic(w, no)
		if err != nil {
			log.Fatalf("コミック情報のインデクスに失敗しました (No.%d): %s", no, err)
			continue
		}
		sep = ","
	}
	fmt.Fprint(w, "]")
}

func addIndexOfAComic(w io.Writer, no int) (*Comic, error) {
	comic, err := downloadComic(no)
	if err != nil {
		return nil, err
	}

	err = writeIndex(w, comic)
	if err != nil {
		return nil, err
	}

	return comic, nil
}

func downloadComic(no int) (*Comic, error) {
	var url string
	if no > 0 {
		url = fmt.Sprintf("https://xkcd.com/%d/info.0.json", no)
	} else {
		url = "https://xkcd.com/info.0.json"
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 5.8 節「遅延関数呼び出し」参照

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}

	return &comic, nil
}

func writeIndex(w io.Writer, comic *Comic) error {
	index, err := json.Marshal(comic)
	if err != nil {
		return err
	}
	fmt.Fprint(w, string(index))
	return nil
}
