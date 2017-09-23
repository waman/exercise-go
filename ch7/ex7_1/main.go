// 【練習問題 7.1】
// ByteCounter の考え方を利用して、ワードと行に対するカウンターを実装しなさい。
// bufio.ScanWords が役に立つでしょう。
package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"io/ioutil"
	"bufio"
	"bytes"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c++
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanLines)
	for s.Scan() {
		*c++
	}
	return len(p), nil
}

// 実行例：
//
//   > go run ./ch7/ex7_1/main.go ./resources/TheGoBlog-strings.txt
func main() {
	var r io.Reader
	if len(os.Args) == 1 {
		r = os.Stdin
	}else{
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		r = f
	}

	content, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	// WordCounter
	var wc WordCounter
	wc.Write(content)
	fmt.Printf("単語数： %d\n", wc)

	// LineCounter
	var lc LineCounter
	lc.Write(content)
	fmt.Printf("行数　： %d\n", lc)
}
